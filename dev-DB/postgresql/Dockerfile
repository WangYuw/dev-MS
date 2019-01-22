FROM postgres
LABEL maintainer="yuwei"

WORKDIR /docker-entrypoint-initdb.d
COPY init.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_DB registry

EXPOSE 5432