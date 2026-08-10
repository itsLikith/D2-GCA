FROM golang:1.26.3-alpine AS builder

WORKDIR /app

ENV GOWORK=off

# Copy shared core packages and simulation service code
COPY pkg/ pkg/
COPY services/simulation-service/ services/simulation-service/

# Compile the simulation service binary
WORKDIR /app/services/simulation-service
RUN go build -o /app/simulation-service ./cmd/server/main.go

# Minimal production image
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/simulation-service .

EXPOSE 8082

CMD ["./simulation-service"]
