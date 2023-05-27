#!/bin/sh

set -e

# had to prepend DB_SOURCE for golang migrate to run properly
echo "Run migrations"
/app/migrate -path /app/migration -database "mysql://$DB_SOURCE" -verbose up

echo "Run the app"
exec "$@"
