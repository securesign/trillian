FROM postgres:17.11@sha256:e38411452a464af89e5adadb8d223bf53b898d47d6ef918b2d58c08707350449

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
