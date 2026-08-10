FROM golang:1.26.3-alpine AS builder

WORKDIR /app

ENV GOWORK=off

# Copy shared core packages and DME service code
COPY pkg/ pkg/
COPY services/dme-service/ services/dme-service/

# Compile the DME service binary
WORKDIR /app/services/dme-service
RUN go build -o /app/dme-service ./cmd/server/main.go

# Minimal production image
FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/dme-service .

EXPOSE 8081

CMD ["./dme-service"]