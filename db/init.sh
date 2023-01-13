#!/bin/bash

echo "*** Starting Migrations for PostgreSQL Database ***"
migrate -path="/migrations/" -database="postgres://postgres@/loonie?host=/var/run/postgresql" up
echo "*** Done ***"