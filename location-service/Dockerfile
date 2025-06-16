FROM golang:1.24.9-alpine3.20 AS builder

WORKDIR /app
COPY . .

COPY go.mod go.sum ./
RUN go mod tidy

RUN go build -o main ./main.go

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 5006 3333

CMD ["/app/main", "serve-rest"]
