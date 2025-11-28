FROM postgres:17.7@sha256:f337b026cb1fb954a93f2e33a62bf9ea7fffc1dafc1586c53b93922a8d6ee018

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
