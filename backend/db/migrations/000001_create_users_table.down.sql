CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "email" VARCHAR(256) UNIQUE NOT NULL,
  "password" VARCHAR(256) NOT NULL,
  "created_at" TIMESTAMPZ NOT NULL DEFAULT (NOW()),
  "updated_at" TIMESTAMPZ NOT NULL DEFAULT (NOW())
);