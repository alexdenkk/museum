FROM golang:1.19

LABEL maintaner="alexdenkk@yandex.ru"

WORKDIR server/

COPY . .

RUN go mod download

COPY go.mod go.sum ./

RUN go build -o server cmd/main.go
