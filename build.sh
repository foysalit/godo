#!/bin/bash
rm ./godo
go build

echo "Build Complete. Firing up server...."

export DB_ADDR="root:admin@/go_do"

./godo