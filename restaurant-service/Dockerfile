FROM golang:1.23.2-alpine3.20 AS builder

WORKDIR /app
COPY . .

COPY go.mod go.sum ./

RUN go mod tidy

# Build the Go application
RUN go build -o main ./main.go

FROM alpine:3.14

WORKDIR /app

# Copy the compiled binary and configuration file from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/config.json .
COPY --from=builder /app/migrations /app/migrations

# Expose the required ports
EXPOSE 6001 3345

# Run the application with the configuration file
CMD ["/app/main", "-c", "./config.json"]  
