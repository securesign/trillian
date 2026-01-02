FROM mysql:9.5@sha256:5ca0a273ed28c73acaef91da8bf1eca3711bee94bce4c378d42846375e645a72


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
