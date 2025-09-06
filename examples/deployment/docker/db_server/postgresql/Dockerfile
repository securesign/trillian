FROM postgres:17.6@sha256:d17be73073d6e73a005c239c67b8c9f0331925d05bd874964f445a9f698fbf72

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
