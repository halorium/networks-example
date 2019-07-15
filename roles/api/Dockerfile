FROM scratch
MAINTAINER Loren Hale <loren.hale@gmail.com>

COPY roles-api /
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY tls/files/* /tls/files/

EXPOSE 80 443
EXPOSE 81 444

ENTRYPOINT ["/roles-api"]
