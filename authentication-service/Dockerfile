# base go image
FROM golang:1.20-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o authApp ./cmd/api

RUN chmod +x /app/authApp


# build a tiny Dockerimage
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/authApp /app

CMD ["/app/authApp"]
