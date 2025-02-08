-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  public.product (
    id VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    price DECIMAL NOT NULL,
    tax DECIMAL NOT NULL
  );

ALTER TABLE
  public.product
ADD
  CONSTRAINT product_pkey PRIMARY KEY (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.product;
-- +goose StatementEnd
