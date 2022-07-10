FROM postgres

ENV POSTGRES_DB numberplates

COPY createdb.sql /docker-entrypoint-initdb.d/
