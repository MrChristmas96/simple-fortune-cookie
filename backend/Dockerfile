# Use the official Go image as the base stage
FROM golang:1.17 AS build-backend

# Set the working directory inside the container
WORKDIR /src

# Copy the module files first to leverage Docker's caching mechanism
COPY ./go.mod ./go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY ./ ./

# Build the Go app for an alpine-linux target
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend-app .

### Start a new stage from scratch
FROM alpine:latest AS runtime-backend

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build-backend /src/backend-app .

# Expose the backend's port
EXPOSE 9000

# Command to run the application
CMD ["./backend-app"]
