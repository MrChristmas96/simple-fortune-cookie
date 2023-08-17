# Use the official Go image as the base stage
FROM golang:1.17 AS build-frontend

# Set the working directory inside the container
WORKDIR /src

# Copy the module files first to leverage Docker's caching mechanism
COPY ./go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY ./ ./

# Build the Go app for an alpine-linux target
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o frontend-app .

### Start a new stage from scratch
FROM alpine:latest AS runtime-frontend

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build-frontend /src/frontend-app .

# Expose the frontend's port
EXPOSE 8080

# Command to run the application
CMD ["./frontend-app"]
