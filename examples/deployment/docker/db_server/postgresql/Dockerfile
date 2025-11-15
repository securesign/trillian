FROM postgres:17.7@sha256:2b6239f32680c308921cf212902f881355c31ae92da0795b4fef29d8f0033021

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
