FROM postgres:17.6@sha256:0b6428e8c09651398137d2b3308a6ad87e73ac15fc38729891c16d942e947d3d

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
