FROM golang:1.8-alpine

ADD . /go/src/gitlab.com/ballad89/location-service

ADD GeoLite2-Country.mmdb /go/src/gitlab.com/ballad89/GeoLite2-Country.mmdb

WORKDIR /go/src/gitlab.com/ballad89/location-service

RUN go install .

ENTRYPOINT location-service

EXPOSE 1989
