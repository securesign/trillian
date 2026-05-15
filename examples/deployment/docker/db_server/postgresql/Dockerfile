FROM postgres:17.10@sha256:2293160ed74e619c591f6c52f956eefe2092d07d440a0bbb518753973d612af7

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
