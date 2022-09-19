FROM golang:1.18-alpine as builder

WORKDIR /app

COPY . /app
RUN go mod download && go mod verify 
RUN go install github.com/cosmtrek/air@latest
COPY . .

CMD [ "air" ]