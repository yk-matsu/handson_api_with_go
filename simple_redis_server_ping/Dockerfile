FROM golang:1.13.8-stretch

WORKDIR /go/src/app
COPY . .

RUN go build

CMD ["go", "run", "main.go"]