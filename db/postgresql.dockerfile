FROM postgres

ENV POSTGRES_DB numberplates

COPY europa.csv /tmp/
RUN chmod a+r /tmp/europa.csv
COPY createdb.sql /docker-entrypoint-initdb.d/
