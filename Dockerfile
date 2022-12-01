FROM golang:alpine

WORKDIR /api

COPY main.go .

COPY go.mod .

RUN  go build

CMD ["./api"]
