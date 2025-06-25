-- +migrate Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- CREATE EXTENSION IF NOT EXISTS pg_uuidv7;
CREATE ROLE admin WITH LOGIN PASSWORD 'admin';
CREATE ROLE readonly WITH LOGIN PASSWORD 'readonly';

-- +migrate Down
DROP EXTENSION IF EXISTS "uuid-ossp" CASCADE;
DROP ROLE admin;
DROP ROLE readonly;
