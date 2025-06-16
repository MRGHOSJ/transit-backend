FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o transit-backend ./cmd/server/main.go


# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/transit-backend .
COPY --from=builder /app/data/transport.json ./data/

EXPOSE 8081

CMD ["./transit-backend"]