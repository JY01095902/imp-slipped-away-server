FROM golang:latest

MAINTAINER Shi.Yanxun

RUN go get github.com/gorilla/websocket

WORKDIR /go/src/the-imp-slipped-away-server

ADD . /go/src/the-imp-slipped-away-server

RUN go install

EXPOSE 9000

CMD ["/go/bin/the-imp-slipped-away-server"]