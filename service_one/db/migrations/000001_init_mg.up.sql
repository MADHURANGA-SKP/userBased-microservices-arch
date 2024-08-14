CREATE TABLE "users" (
    "user_id" bigserial PRIMARY KEY,
    "user_name" varchar NOT NULL,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
    "is_email_verified" boolean NOT NULL DEFAULT false,
    "role" varchar NOT NULL DEFAULT 'customer',
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "orders" (
    "order_id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "item_id" bigint NOT NULL,
    "status" varchar NOT NULL,
    "payment_link" varchar,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "items" (
    "item_id" bigserial PRIMARY KEY,
    "user_id" bigint NOT NULL,
    "item_name" varchar NOT NULL,
    "quantity" bigint NOT NULL,
    "price" bigint NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX  ON "users" ("user_id");

CREATE INDEX  ON "orders" ("order_id");

CREATE INDEX  ON "orders" ("user_id");

CREATE INDEX  ON "orders" ("item_id");

CREATE INDEX ON "orders" ("order_id", "user_id", "item_id");

CREATE INDEX  ON "items" ("item_id");

CREATE INDEX  ON "items" ("user_id");

CREATE INDEX ON "items" ("item_id", "user_id" );

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE;

ALTER TABLE "orders" ADD FOREIGN KEY ("item_id") REFERENCES "items" ("item_id") ON DELETE CASCADE;

ALTER TABLE "items" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id") ON DELETE CASCADE;
