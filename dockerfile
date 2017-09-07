FROM golang:latest

MAINTAINER Shi.Yanxun

RUN go get github.com/gorilla/websocket
RUN go get github.com/aliyun/aliyun-tablestore-go-sdk
RUN go get github.com/labstack/echo

WORKDIR /go/src/imp-slipped-away-server

ADD . /go/src/imp-slipped-away-server

RUN go install

EXPOSE 9000

CMD ["/go/bin/imp-slipped-away-server"]