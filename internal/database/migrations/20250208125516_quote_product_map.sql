-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  "public"."quote_product_map" (
    "id" VARCHAR(255) NOT NULL,
    "quote_fk" VARCHAR(255),
    "product_fk" VARCHAR(255),
    CONSTRAINT "quote_product_map_pkey" PRIMARY KEY ("id"),
    FOREIGN KEY ("quote_fk") REFERENCES "public"."quote"("id"),
    FOREIGN KEY ("product_fk") REFERENCES "public"."product"("id")
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "public"."quote_product_map"; 
-- +goose StatementEnd
