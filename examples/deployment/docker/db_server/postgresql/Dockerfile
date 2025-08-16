FROM postgres:17.6@sha256:a7fdfd815005ea658281c60b06020f49ab8a6bbe21434d3a59fe14654f6d4293

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
