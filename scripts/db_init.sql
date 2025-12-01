CREATE EXTENSION IF NOT EXISTS citext;
-- 1. Create the database (run once manually, NOT in migrate)
CREATE DATABASE social;

-- 2. Connect to the new database
\c social

-- 3. Enable CITEXT inside the database
