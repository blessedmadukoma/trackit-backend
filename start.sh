#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
echo "Finished migration on $DB_SOURCE"
echo "Should start the app on $SERVER_ADDRESS"
echo "starting the app"
echo "$@" # take all parameters and run in this script