FROM golang:1.23.0-alpine

WORKDIR /app


COPY . .

RUN go mod download
RUN go mod tidy