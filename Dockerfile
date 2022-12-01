FROM golang:alpine

WORKDIR /api

COPY main.go .

COPY go.mod .

COPY go.sum .

RUN  go build

CMD ["./api"]
