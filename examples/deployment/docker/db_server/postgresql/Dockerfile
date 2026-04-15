FROM postgres:17.9@sha256:7ad98329d513dd497293b951c195ca354274a77f12ddbbbbf85e68a811823d72

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
