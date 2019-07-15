FROM scratch
MAINTAINER Loren Hale <loren.hale@gmail.com>

COPY networks-worker /
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY tls/files/* /tls/files/

ENTRYPOINT ["/networks-worker"]
