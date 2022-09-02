FROM golang:1.18-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o brokerApp ./cmd/api

CMD ["/app/brokerApp"]