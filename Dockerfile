FROM golang:1.22

WORKDIR /go/src/app

COPY . .

RUN go build -o main src/main.go

CMD ["./main"]