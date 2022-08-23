CREATE TABLE "applications" (
  "id" uuid NOT NULL PRIMARY KEY,
  "first_name" VARCHAR NOT NULL,
  "last_name" VARCHAR NOT NULL,
  "status" VARCHAR NOT NULL DEFAULT 'pending'
);
