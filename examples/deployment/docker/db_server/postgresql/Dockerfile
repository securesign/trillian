FROM postgres:17.5@sha256:fa6e5715d6521a4a9d4cc47015b6eb8fc4f2abaae4bc015be7a5209328ce196c

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
