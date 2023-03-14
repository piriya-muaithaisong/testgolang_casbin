CREATE TABLE "users" (
  "id" uuid UNIQUE,
  "username" varchar UNIQUE,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "role" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id", "username")
);

CREATE TABLE "assets" (
  "id" uuid UNIQUE PRIMARY KEY,
  "asset_name" varchar NOT NULL,
  "owner" uuid NOT NULL,
  "tester" uuid NOT NULL
);

CREATE TABLE "result" (
  "id" uuid UNIQUE PRIMARY KEY,
  "result_name" varchar NOT NULL,
  "assets" uuid NOT NULL,
  "tester" uuid NOT NULL
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL,
  "expires_at" timestamptz NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "assets" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");

ALTER TABLE "assets" ADD FOREIGN KEY ("tester") REFERENCES "users" ("id");

ALTER TABLE "result" ADD FOREIGN KEY ("assets") REFERENCES "assets" ("id");

ALTER TABLE "result" ADD FOREIGN KEY ("tester") REFERENCES "users" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
