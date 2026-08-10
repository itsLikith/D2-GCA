FROM golang:1.26.3-alpine AS builder

WORKDIR /app

ENV GOWORK=off

# Copy shared core packages and RNAV service code
COPY pkg/ pkg/
COPY services/rnav-service/ services/rnav-service/

# Compile the RNAV service binary
WORKDIR /app/services/rnav-service
RUN go build -o /app/rnav-service ./cmd/server/main.go

# Minimal production image
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/rnav-service .

EXPOSE 8080

CMD ["./rnav-service"]
