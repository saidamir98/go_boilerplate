CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS  "application" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "body" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);