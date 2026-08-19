# syntax=docker/dockerfile:1

FROM golang:1.25-alpine AS builder

ENV CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /bin/api ./cmd/api

FROM alpine:3.21 AS final

RUN apk add --no-cache ca-certificates wget

WORKDIR /app
COPY --from=builder /bin/api /bin/api

EXPOSE 8080

ENTRYPOINT ["/bin/api"]
