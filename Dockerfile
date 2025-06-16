FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o transit-backend ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/transit-backend .
COPY data/transport.json ./data/transport.json
CMD ["./transit-backend"]