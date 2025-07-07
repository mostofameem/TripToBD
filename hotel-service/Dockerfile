FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /app
COPY . .

COPY go.mod go.sum ./
RUN go mod tidy
# Build the Go application
RUN go build -o main ./main.go

# Stage 2: Create a minimal runtime image
FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/migrations /app/migrations

# Expose the required ports
EXPOSE 5005 3355

# Command to run the executable
CMD ["/app/main", "serve-rest"]