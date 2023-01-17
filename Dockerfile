FROM golang:1.19-alpine

WORKDIR /app

COPY . /app

RUN go mod download

EXPOSE 9000:9000

CMD go run .
