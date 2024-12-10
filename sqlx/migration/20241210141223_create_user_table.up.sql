-- Enable pgcrypto for cryptographic functions
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Create ULID generation function
CREATE OR REPLACE FUNCTION ulid_generate() RETURNS text AS $$
DECLARE
    timestamp  bigint;
    entropy    bytea;
    ulid      text;
BEGIN
    -- Get current timestamp in milliseconds
    timestamp := (EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) * 1000)::bigint;
    
    -- Generate 10 bytes of entropy
    entropy := gen_random_bytes(10);
    
    -- Convert timestamp to base32
    ulid := lpad(to_hex(timestamp), 12, '0');
    
    -- Convert entropy to base32 and append
    ulid := ulid || encode(entropy, 'hex');
    
    RETURN ulid;
END;
$$ LANGUAGE plpgsql;

-- Create users table
CREATE TABLE IF NOT EXISTS "users" (
    id text PRIMARY KEY DEFAULT ulid_generate(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);