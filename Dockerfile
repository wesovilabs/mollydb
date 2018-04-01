FROM alpine:3.6
ADD ./build/mollydb.linux /usr/local/bin/mollydb
ADD ./resources/config /var/mollydb/config
ADD ./resources/graphiql /var/mollydb/graphiql

VOLUME ["/var/mollydb/config"]

EXPOSE 9090
ENTRYPOINT /usr/local/bin/mollydb -config=/var/mollydb/config/mollydb.json