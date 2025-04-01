#!/bin/bash

set -e
set -u

sed "s/\${USER}/$POSTGRES_USER/g; s/\${PWD}/$POSTGRES_PASSWORD/g" /queries/CREATE_USER.sql | psql -U postgres -d postgres
sed "s/\${NAME}/$POSTGRES_DB/g; s/\${USER}/$POSTGRES_USER/g" /queries/CREATE_DATABASE.sql | psql -U $POSTGRES_USER -d postgres

# TYPES
psql -U $POSTGRES_USER -d $POSTGRES_DB -f /queries/CREATE_USER_GENDER_ENUM.sql
psql -U $POSTGRES_USER -d $POSTGRES_DB -f /queries/CREATE_USER_STATUS_ENUM.sql