-- Create "organizations" table
CREATE TABLE "organizations" ("id" uuid NOT NULL, "name" character varying NOT NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, "public_id" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create index "organizations_public_id_key" to table: "organizations"
CREATE UNIQUE INDEX "organizations_public_id_key" ON "organizations" ("public_id");
-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "email" character varying NOT NULL, "name" character varying NOT NULL, "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
-- Create "organization_users" table
CREATE TABLE "organization_users" ("organization_id" uuid NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("organization_id", "user_id"), CONSTRAINT "organization_users_organization_id" FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON DELETE CASCADE, CONSTRAINT "organization_users_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE);
