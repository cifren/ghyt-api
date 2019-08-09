FROM golang:1.12.7

WORKDIR /go
COPY . .

RUN go get -d -v ./...
RUN go build main.go

CMD ["./main"]