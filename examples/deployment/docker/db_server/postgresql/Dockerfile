FROM postgres:17.6@sha256:b480430782a9bd1c8a6835fb5b70f89f34a70132c2f6182e534f65688bce063b

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
