FROM golang:1.20.0-alpine AS build

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x /app/brokerApp

FROM alpine:latest

RUN mkdir /app

COPY --from=build /app/brokerApp /app

WORKDIR /app

CMD ["./brokerApp"]