FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o auditservice ./cmd/main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/auditservice /app/auditservice

EXPOSE 50051

CMD ["/app/auditservice"]