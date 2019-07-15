FROM scratch
MAINTAINER Loren Hale <loren.hale@gmail.com>

COPY cycler-service /
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/cycler-service"]
