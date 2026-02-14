FROM postgres:17.8@sha256:1eada89ec2521ae0f61125f6bdaf83bbbf88a43fecda48d112a83bb92df75c36

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY storage/postgresql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d
