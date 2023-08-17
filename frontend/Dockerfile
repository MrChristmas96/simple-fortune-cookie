# Use the official Go image as the base stage
FROM golang:1.21 AS build-frontend
#Build Go 
WORKDIR /src/
COPY ./go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o frontend-app .

#Build Docker
FROM alpine:latest AS runtime-frontend
WORKDIR /app
COPY --from=build-frontend /src/ .
EXPOSE 8080
CMD ["./frontend-app"]
