FROM golang:1.26.3-alpine AS builder

WORKDIR /app

ENV GOWORK=off

# Copy shared core packages and gateway service code
COPY pkg/ pkg/
COPY services/gateway/ services/gateway/

# Compile the gateway service binary
WORKDIR /app/services/gateway
RUN go build -o /app/gateway-service ./cmd/server/main.go

# Minimal production image
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/gateway-service .

EXPOSE 8000

CMD ["./gateway-service"]