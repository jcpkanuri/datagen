#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
  create table department(
    id smallserial primary key,
    name varchar(10) not null
  );

  insert into department (name) values('presales');
  insert into department (name) values('design');
  insert into department (name) values('sales');
  insert into department (name) values('account');
  insert into department (name) values('quality');
  insert into department (name) values('install');
  insert into department (name) values('service');
  insert into department (name) values('audit');
  COMMIT;
EOSQL