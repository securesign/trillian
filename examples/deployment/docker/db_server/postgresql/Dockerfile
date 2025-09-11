FROM postgres:17.6@sha256:feff5b24fedd610975a1f5e743c51a4b360437f4dc3a11acf740dcd708f413f6

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
