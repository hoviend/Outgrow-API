-- Create index "organizationeventtype_organization_id_name" to table: "organization_event_types"
DROP INDEX IF EXISTS "organization_event_types_name_key";
CREATE UNIQUE INDEX "organizationeventtype_organization_id_name" ON "organization_event_types" ("organization_id", "name");

-- Create index "organizationaccounttype_organization_id_name" to table: "organization_account_types"
DROP INDEX IF EXISTS "organization_account_types_name_key";
CREATE UNIQUE INDEX "organizationaccounttype_organization_id_name" ON "organization_account_types" ("organization_id", "name");

