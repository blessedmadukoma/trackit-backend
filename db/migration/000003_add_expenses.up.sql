CREATE TABLE "expenses" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "userid" bigint NOT NULL,
  "email" varchar NOT NULL,
  "amount" varchar NOT NULL,
  "description" varchar NOT NULL,
  "tag" varchar NOT NULL,
  "date" timestamptz NOT NULL DEFAULT (now()),
  -- "budgetid" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

ALTER TABLE "expenses" ADD FOREIGN KEY ("userid") REFERENCES "user" ("id");