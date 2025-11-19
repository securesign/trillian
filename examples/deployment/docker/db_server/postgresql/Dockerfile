FROM postgres:17.7@sha256:5ecd7c63c618bed8b554567d3e9959636956be2058d8acbf12d0313dd16f67f8

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
