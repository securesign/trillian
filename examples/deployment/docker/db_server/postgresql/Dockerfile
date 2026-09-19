FROM postgres:18.6@sha256:10037cee1d5ebfc057cd71b3365466cf27ce8d4bd0c4f409e426fcb2b3fb12e7

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
