#bin/bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	ALTER USER postgres PASSWORD 'postgres';
EOSQL

createdb registry
RUN psql -U postgres -f init.sql -d registry