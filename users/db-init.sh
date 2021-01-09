#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER temp WITH password 'qwerty123';
    CREATE DATABASE users_db;
    GRANT ALL PRIVILEGES ON DATABASE users_db TO temp;
EOSQL