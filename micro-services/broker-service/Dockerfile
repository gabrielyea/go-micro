FROM golang:1.18-alpine as builder

WORKDIR /app

COPY . /app
RUN go mod download && go mod verify 
RUN go install github.com/cosmtrek/air@latest
COPY . .

CMD [ "air" ]

# build binary
# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api
# RUN chmod +x /app/brokerApp

# # build a tiny docker image
# FROM alpine:latest
# RUN mkdir /app

# COPY --from=builder /app/brokerApp /app
# CMD ["air"]