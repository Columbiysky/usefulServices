-- Create "accounts" table
CREATE TABLE "public"."accounts" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "login" text NOT NULL,
  "password" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_accounts_deleted_at" to table: "accounts"
CREATE INDEX "idx_accounts_deleted_at" ON "public"."accounts" ("deleted_at");
-- Create "tokens" table
CREATE TABLE "public"."tokens" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "token_value" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_tokens_deleted_at" to table: "tokens"
CREATE INDEX "idx_tokens_deleted_at" ON "public"."tokens" ("deleted_at");
-- Create "account_tokens" table
CREATE TABLE "public"."account_tokens" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "account_refer" bigint NULL,
  "token_refer" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_account_tokens_account" FOREIGN KEY ("account_refer") REFERENCES "public"."accounts" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_account_tokens_token" FOREIGN KEY ("token_refer") REFERENCES "public"."tokens" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_account_tokens_deleted_at" to table: "account_tokens"
CREATE INDEX "idx_account_tokens_deleted_at" ON "public"."account_tokens" ("deleted_at");
