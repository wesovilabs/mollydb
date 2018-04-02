#!/bin/bash

curl -XPOST http://mollydb:9090/graphql \
-H 'Content-Type: application/graphql' \
-d "mutation ms {
        register (path: \"/var/mollydb/storage/ms\", name: \"ms\")
        {name}
    }" \
-i

