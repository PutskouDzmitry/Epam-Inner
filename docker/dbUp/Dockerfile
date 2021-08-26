FROM timescale/timescaledb:latest-pg13

COPY ./pkg/db/script/init.sql /docker-entrypoint-initdb.d/init.sql
