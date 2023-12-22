FROM golang:1.21 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o bin/server


FROM scratch
WORKDIR /app

VOLUME /app/data

COPY --from=builder /build/bin/ .

COPY migration.sql .

ENV DB_PATH=/app/data/todo.db
ENV APP_INTERFACE=:8080

EXPOSE 8080

CMD ["./server"]
