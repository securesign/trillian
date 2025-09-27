FROM postgres:17.6@sha256:cecd364cc7ec03ab6148ab650efc5200548a89fd2e4f47a8ba6e6cf08c18805d

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
