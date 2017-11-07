FROM golang

MAINTAINER phanvudinh <itphanvudinh@gmail.com>

EXPOSE 8080

VOLUME /go/src/github.com/phanvudinh/go-practice-app

CMD cd /go/src/github.com/phanvudinh/go-practice-app && go get && go build && cd ../../../../bin && ./go-practice-app
