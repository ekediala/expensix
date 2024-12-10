-- write down migration here

DROP TABLE IF EXISTS "users";

DROP FUNCTION IF EXISTS ulid_generate;

DROP EXTENSION IF EXISTS "pgcrypto";