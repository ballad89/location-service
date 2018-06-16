FROM golang:1.10.3-alpine

ADD . /go/src/gitlab.com/ballad89/location-service

ADD http://geolite.maxmind.com/download/geoip/database/GeoLite2-Country.tar.gz /go/src/gitlab.com/ballad89/location-service/

WORKDIR /go/src/gitlab.com/ballad89/location-service

RUN tar -xzf GeoLite2-Country.tar.gz

RUN go build


FROM gliderlabs/alpine:3.6

WORKDIR /root/
COPY --from=0 /go/src/gitlab.com/ballad89/location-service/location-service .
COPY --from=0 /go/src/gitlab.com/ballad89/location-service/GeoLite2-Country_*/GeoLite2-Country.mmdb .

ENV PORT ":3000"
CMD ["./location-service"]

EXPOSE 3000
