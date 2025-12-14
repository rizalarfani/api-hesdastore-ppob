# BUILDER
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata openssh

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ppob-service

# RUNNER
FROM alpine:3.18

RUN apk add --no-cache ca-certificates tzdata openssh

WORKDIR /app

COPY --from=builder /app/ppob-service .

EXPOSE 8001

CMD ["./ppob-service"]