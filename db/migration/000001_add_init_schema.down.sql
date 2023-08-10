-- ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";
-- ALTER TABLE "product" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");
-- ALTER TABLE "product" ADD FOREIGN KEY ("ingredients_id") REFERENCES "ingredients" ("id");
-- ALTER TABLE "product" ADD FOREIGN KEY ("picture_id") REFERENCES "picture" ("id");
-- ALTER TABLE "ingredients" ADD FOREIGN KEY ("picture_id") REFERENCES "picture" ("id");
-- ALTER TABLE "feeling" ADD FOREIGN KEY ("product_id") REFERENCES "product" ("id");

ALTER TABLE  "product" DROP CONSTRAINT  "product_category_id_fkey";
ALTER TABLE  "product" DROP CONSTRAINT  "product_ingredients_id_fkey";
ALTER TABLE  "product" DROP CONSTRAINT  "product_picture_id_fkey";
ALTER TABLE  "ingredients" DROP CONSTRAINT  "ingredients_picture_id_fkey";
ALTER TABLE  "feeling" DROP CONSTRAINT  "feeling_product_id_fkey";
DROP TABLE IF EXISTS ingredients;
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS picture;
DROP TABLE IF EXISTS category;
DROP TABLE IF EXISTS feeling;
