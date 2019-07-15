FROM golang:1.9.2-alpine
MAINTAINER Loren Hale <loren.hale@gmail.com>

ENV IN_MIGRATIONS_PATH="/migrations"

RUN apk update                && \
    apk add bash              && \
    apk add postgresql-client && \
    apk add curl

COPY migrator-service /
COPY executables/* /migrations/

WORKDIR /

CMD ["/migrator-service"]
