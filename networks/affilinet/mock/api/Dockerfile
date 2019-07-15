FROM scratch
MAINTAINER Loren Hale <loren.hale@gmail.com>

COPY networks-affilinet-mock-api /
COPY response-authentication.xml /
COPY response-page-1.xml /
COPY response-page-2.xml /

EXPOSE 80

ENTRYPOINT ["/networks-affilinet-mock-api"]
