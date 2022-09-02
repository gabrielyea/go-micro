FROM golang:1.18-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o authApp ./cmd/api

CMD ["/app/authApp"]