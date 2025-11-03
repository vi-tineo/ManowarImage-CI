# Etapa 1: build
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY ./main.go .
RUN go mod init manowar
RUN go build -o server .

# Etapa 2: runtime
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]
