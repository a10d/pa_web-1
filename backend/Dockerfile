FROM golang:1.21 AS builder

WORKDIR /build

# Download the go modules
COPY go.mod go.sum ./
RUN go mod download


COPY . .

ARG GOOS=linux
ARG GOARCH=amd64
ARG CGO_ENABLED=0

# Build the binary
RUN go build -o bin/server

FROM scratch
WORKDIR /app

VOLUME /app/data

# Copy the binary from the builder stage
COPY --from=builder /build/bin/ .

# Copy all locally required files
COPY migration.sql .

# Set the environment variables
ENV DB_PATH=/app/data/todo.db
ENV APP_INTERFACE=:8080

# Expose the port
EXPOSE 8080

CMD ["./server"]
