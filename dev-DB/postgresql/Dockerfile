FROM postgres
LABEL maintainer="yuwei"

USER postgres

WORKDIR /sql
COPY init.sql /sql

RUN echo "host all  all    md5" >> /etc/postgresql/pg_hba.conf
RUN echo "listen_addresses='*'" >> /etc/postgresql/postgresql.conf

RUN service postgresql restart

#RUN psql --command "ALTER USER postgres WITH ENCRYPTED PASSWORD 'postgres';"
#RUN createdb registry
#RUN psql -U postgres -f init.sql -d registry

EXPOSE 5432