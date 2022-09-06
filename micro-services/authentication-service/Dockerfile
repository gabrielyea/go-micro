FROM golang:1.18-alpine

WORKDIR /app

COPY . /app
RUN go mod download && go mod verify 
RUN go install github.com/cosmtrek/air@latest
COPY . .


CMD ["air"]