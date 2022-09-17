FROM golang:1.19-alpine AS build-env

ENV CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64

RUN mkdir /app

ADD ./backend /app/

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]
