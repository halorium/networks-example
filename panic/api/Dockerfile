FROM scratch
MAINTAINER Tom Mornini <tmornini@networksexample.com>

COPY panic-api /
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY tls/files/* /tls/files/

EXPOSE 80 443
EXPOSE 81 444

ENTRYPOINT ["/panic-api"]
