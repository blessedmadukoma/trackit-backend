CREATE TABLE "user" (
    "id" bigserial NOT NULL PRIMARY KEY,
    "firstname" VARCHAR NOT NULL,
    "lastname" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL UNIQUE,
    "password" VARCHAR NOT NULL,
    "phone" VARCHAR NOT NULL,
    "user_type" boolean NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz
);


ALTER TABLE "tablename" ADD FOREIGN KEY ("userid") REFERENCES "user" ("id") ON DELETE SET DEFAULT;