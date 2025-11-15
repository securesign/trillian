FROM postgres:18.1@sha256:41bfa2e5b168fff0847a5286694a86cff102bdc4d59719869f6b117bb30b3a24

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
