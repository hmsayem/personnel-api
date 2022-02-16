FROM golang:latest

LABEL maintainer="Hossain Mahmud <hmsayem@gmail.com"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV SERVER_PORT=:8000

ENV GOOGLE_APPLICATION_CREDENTIALS=/run/secrets/server-key.json

RUN go build -o rest-server

CMD ["./rest-server"]
