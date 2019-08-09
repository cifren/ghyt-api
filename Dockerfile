FROM golang:1.12.7

WORKDIR /go
COPY . .

RUN go get -d -v main
RUN go install main

CMD ["main"]