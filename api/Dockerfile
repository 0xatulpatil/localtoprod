# Stage 1: Builder
FROM golang:1.24.2-alpine AS builder

RUN apk add --no-cache git gcc musl-dev sqlite-dev

WORKDIR /app

ENV CGO_ENABLED=1
ENV GOOS=linux

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app .

# Stage 2: Minimal runtime
FROM alpine:latest

RUN apk add --no-cache ca-certificates sqlite-libs
RUN apk add --no-cache curl

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/app .

EXPOSE 8080

# Entrypoint
CMD ["./app"]