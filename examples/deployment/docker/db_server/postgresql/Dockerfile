FROM postgres:17.9@sha256:9d33475e46c9c6c09d5d8ee6fe609d05f5351fd44d1079d1d870bd5c404e47c0

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
