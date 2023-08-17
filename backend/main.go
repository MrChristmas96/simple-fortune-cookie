package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"sync"
)

var (
	listFortuneRe   = regexp.MustCompile(`^/fortunes[/]*$`)
	getFortuneRe    = regexp.MustCompile(`^/fortunes[/](\d+)$`)
	randomFortuneRe = regexp.MustCompile(`^/fortunes[/]random$`)
	createFortuneRe = regexp.MustCompile(`^/fortunes[/]*$`)
)

type fortune struct {
	ID      string `json:"id" redis:"id"`
	Message string `json:"message" redis:"message"`
}

type datastore struct {
	m map[string]fortune
	*sync.RWMutex
}

var datastoreDefault = datastore{m: map[string]fortune{
	"1": {ID: "1", Message: "A new voyage will fill your life with untold memories."},
	"2": {ID: "2", Message: "The measure of time to your next goal is the measure of your discipline."},
	"3": {ID: "3", Message: "The only way to do well is to do better each day."},
	"4": {ID: "4", Message: "It ain't over till it's EOF."},
}, RWMutex: &sync.RWMutex{}}

type fortuneHandler struct {
	store *datastore
}

// The simplified error response function
func respondWithError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	_, _ = w.Write([]byte(message))
}

func (h *fortuneHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listFortuneRe.MatchString(r.URL.Path):
		h.List(w, r)
	case r.Method == http.MethodGet && getFortuneRe.MatchString(r.URL.Path):
		h.Get(w, r)
	case r.Method == http.MethodGet && randomFortuneRe.MatchString(r.URL.Path):
		h.Random(w, r)
	case r.Method == http.MethodPost && createFortuneRe.MatchString(r.URL.Path):
		h.Create(w, r)
	default:
		respondWithError(w, http.StatusNotFound, "not found")
	}
}

// Simplified function for writing JSON to the response writer
func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	_, _ = w.Write(jsonBytes)
}

func (h *fortuneHandler) List(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	fortunes := make([]fortune, 0, len(h.store.m))
	for _, v := range h.store.m {
		fortunes = append(fortunes, v)
	}
	h.store.RUnlock()

	writeJSONResponse(w, fortunes)
}

func (h *fortuneHandler) Random(w http.ResponseWriter, r *http.Request) {
	h.store.RLock()
	fortunes := make([]fortune, 0, len(h.store.m))
	for _, v := range h.store.m {
		fortunes = append(fortunes, v)
	}
	h.store.RUnlock()

	if len(fortunes) > 0 {
		u := fortunes[rand.Intn(len(fortunes))]
		r.URL.Path = "/fortunes/" + u.ID
	} else {
		r.URL.Path = "/fortunes/zero"
	}

	h.Get(w, r)
}

func (h *fortuneHandler) Get(w http.ResponseWriter, r *http.Request) {
	matches := getFortuneRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		respondWithError(w, http.StatusNotFound, "not found")
		return
	}

	key := matches[1]

	h.store.RLock()
	u, ok := h.store.m[key]
	h.store.RUnlock()

	if !ok {
		respondWithError(w, http.StatusNotFound, "fortune not found")
		return
	}
	writeJSONResponse(w, u)
}

func (h *fortuneHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u fortune
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	h.store.Lock()
	h.store.m[u.ID] = u
	h.store.Unlock()

	writeJSONResponse(w, u)
}

func main() {
	mux := http.NewServeMux()
	fortuneH := &fortuneHandler{
		store: &datastoreDefault,
	}
	mux.Handle("/fortunes", fortuneH)
	mux.Handle("/fortunes/", fortuneH)

	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
