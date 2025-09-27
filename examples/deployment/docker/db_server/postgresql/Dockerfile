FROM postgres:18.0@sha256:1d288494853e244e7a78d87b3526e650e5221c622f9768ecac9313d0874a9c39

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
