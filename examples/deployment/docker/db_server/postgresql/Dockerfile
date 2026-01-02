FROM postgres:17.7@sha256:bac8128cd62a35471a1e97966861aa2ac88c8d2f408767a425ae70695a607c99

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
