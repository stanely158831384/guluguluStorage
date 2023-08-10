CREATE TABLE "category" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "product" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "category_id" bigint NOT NULL,
  "ingredients_id" bigint NOT NULL,
  "risk_level" smallserial,
  "picture_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ingredients" (
  "id" bigserial PRIMARY KEY,
  "ingredient" varchar[],
  "picture_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "feeling" (
  "id" bigserial PRIMARY KEY,
  "product_id" bigint NOT NULL,
  "user_id" bigserial NOT NULL,
  "username" varchar NOT NULL,
  "comment" varchar NOT NULL,
  "recommend" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "picture" (
  "id" bigserial PRIMARY KEY,
  "link" varchar,
  "user_id" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "product" ("name");

CREATE INDEX ON "product" ("risk_level");

CREATE INDEX ON "product" ("name", "risk_level");

CREATE INDEX ON "feeling" ("username");

CREATE INDEX ON "feeling" ("comment");

CREATE INDEX ON "feeling" ("recommend");

CREATE INDEX ON "picture" ("link");

ALTER TABLE "product" ADD CONSTRAINT "product_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "product" ADD CONSTRAINT "product_ingredients_id_fkey" FOREIGN KEY ("ingredients_id") REFERENCES "ingredients" ("id");

ALTER TABLE "product" ADD CONSTRAINT "product_picture_id_fkey" FOREIGN KEY ("picture_id") REFERENCES "picture" ("id");

ALTER TABLE "ingredients" ADD CONSTRAINT "ingredients_picture_id_fkey" FOREIGN KEY ("picture_id") REFERENCES "picture" ("id");

ALTER TABLE "feeling" ADD CONSTRAINT "feeling_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "product" ("id");

