FROM golang:1.19-alpine

WORKDIR /app

COPY . /app

RUN go mod download
