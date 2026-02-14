FROM mysql:9.6@sha256:db32c8ec843c042a728efb0ac7aa814d6f010eaac8923e20ae0a849d09c5baf8


# TODO(roger2hk): Uncomment the below OS-level packages patch command as this is a temporary workaround to bypass the mysql8 gpg key rotation issue.
# # Patch the OS-level packages and remove unneeded dependencies.
# ENV DEBIAN_FRONTEND=noninteractive
# RUN apt-get update \
#     && apt-get install -y procps \
#     && apt-get -y -o Dpkg::Options::="--force-confdef" -o Dpkg::Options::="--force-confold" upgrade \
#     && apt-get -y autoremove \
#     && rm -rf /var/lib/apt/lists/*

# expects the build context to be: $GOPATH/src/github.com/google/trillian
COPY examples/deployment/docker/db_server/mysql.cnf /etc/mysql/conf.d/trillian.cnf
COPY storage/mysql/schema/storage.sql /docker-entrypoint-initdb.d/storage.sql
RUN chmod -R 775 /docker-entrypoint-initdb.d && \
    chmod 644 /etc/mysql/conf.d/trillian.cnf
