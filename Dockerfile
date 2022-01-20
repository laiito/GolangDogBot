# syntax=docker/dockerfile:1

FROM golang:1.17.6-alpine

WORKDIR /app

COPY . /app
RUN go mod download

RUN go build -o /golangdogbot

EXPOSE 80/tcp

CMD [ "/golangdogbot" ]