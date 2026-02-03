FROM postgres:17.7@sha256:4493696c5ba6a9cd4a3303411d5dd6af52cf5e34cdd87ba8ebf26a19735f84d1

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
