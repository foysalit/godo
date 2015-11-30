#!/bin/bash
rm ./godo
go build

echo "Build Complete. Firing up server...."

export DB_USER="root"
export DB_PASS="admin"
export DB_NAME="go_do"

./godo