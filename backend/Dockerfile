FROM golang:1.19-alpine AS build-env

ENV CGO_ENABLED=0\
    GOOS=linux\
    GOARCH=amd64

RUN mkdir /app
RUN mkdir /app/backend

ADD . /app/backend

WORKDIR /app/backend

RUN go build -o main .

CMD ["/app/backend/main"]
