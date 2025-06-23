FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o backend ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/backend .
RUN chmod +x ./backend

EXPOSE 8080

CMD ["./backend"]