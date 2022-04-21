# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

ARG USER
ARG PASSWORD
ARG DATABASE

ENV USER=$USER
ENV PASSWORD=$PASSWORD
ENV DATABASE=$DATABASE

WORKDIR /scaleflix

COPY . .
RUN go mod download

RUN go build ./cmd/server/.

EXPOSE 8080

CMD /scaleflix/server -dsn="host='db' user=$USER password=$PASSWORD dbname=$DATABASE"