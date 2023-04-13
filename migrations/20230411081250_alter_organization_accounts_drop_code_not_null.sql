-- Modify "organization_accounts" table
ALTER TABLE "organization_accounts" 
    ALTER COLUMN "code" DROP NOT NULL;

DROP INDEX organization_accounts_code_key;

CREATE UNIQUE INDEX organization_accounts_code_key
ON "organization_accounts" ("code")
WHERE "code" IS NOT NULL;

-- Modify "events" table
ALTER TABLE "events" 
    ADD COLUMN "organization_id" uuid NOT NULL REFERENCES "organizations" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;


