-- Modify "organization_accounts" table
ALTER TABLE "organization_accounts" 
    ADD COLUMN "code" character varying NOT NULL,
    ADD COLUMN "balance" double precision NOT NULL DEFAULT 0;
-- Create index "organization_accounts_code_key" to table: "organization_accounts"
CREATE UNIQUE INDEX "organization_accounts_code_key" ON "organization_accounts" ("code");

-- Create "events" table
CREATE TABLE "events" (
    "id" uuid NOT NULL,
    "amount" double precision NOT NULL DEFAULT 0,
    "notes" text NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "event_type_id" bigint NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "events_organization_event_types_events" FOREIGN KEY ("event_type_id") REFERENCES "organization_event_types" ("id") ON DELETE RESTRICT ON UPDATE CASCADE);

-- Modify "transactions" table
ALTER TABLE "transactions" 
    ADD COLUMN "event_id" uuid NOT NULL,
    ADD CONSTRAINT "transactions_events_transactions" FOREIGN KEY ("event_id") REFERENCES "events" ("id") ON DELETE RESTRICT ON UPDATE CASCADE;

