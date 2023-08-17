#Build Go
FROM golang:1.21 AS build-backend
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend-app .

#Build Docker
FROM alpine:latest AS runtime-backend
WORKDIR /app
COPY --from=build-backend /src/ .
EXPOSE 9000
CMD ["./backend-app"]
