# STAGE 1

FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o myapp main.go

# STAGE 2

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/myapp .

EXPOSE 8080

CMD ["./myapp"]
