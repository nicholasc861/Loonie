#!/bin/bash

echo "*** Starting Migrations for PostgreSQL Database ***"
migrate -path="/docker-entrypoint-initdb.d/schema/" -database="postgres://postgres@/postgres?host=/var/run/postgresql" up
echo "*** Done ***"