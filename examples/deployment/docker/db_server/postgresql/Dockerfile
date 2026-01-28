FROM postgres:17.7@sha256:7352e0c4d62bbac8aa69d95e40220a60967c4a19f9c4f65b4d118175f7ce9e3b

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
