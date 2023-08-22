-- Create "users" table
CREATE TABLE "users" ("id" bigserial NOT NULL, "created_at" timestamptz NULL, "updated_at" timestamptz NULL, "deleted_at" timestamptz NULL, "name" text NOT NULL, "email" text NOT NULL, "birthday" timestamptz NULL, PRIMARY KEY ("id"), CONSTRAINT "chk_users_email" CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'::text));
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
-- Create index "idx_users_name" to table: "users"
CREATE INDEX "idx_users_name" ON "users" ("name");
