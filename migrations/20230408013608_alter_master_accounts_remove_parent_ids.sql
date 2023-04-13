-- Modify "master_accounts" table
ALTER TABLE "master_accounts" DROP COLUMN parent_id;

-- Modify "organization_accounts" table
ALTER TABLE "organization_accounts" DROP COLUMN parent_id;
