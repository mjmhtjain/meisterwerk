-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  public.order (
    id VARCHAR(255) NOT NULL,
    quote_fk VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
  );

ALTER TABLE
  public.order
ADD
  CONSTRAINT order_pkey PRIMARY KEY (id);

ALTER TABLE
  public.order
ADD
  CONSTRAINT order_quote_fk_fkey FOREIGN KEY (quote_fk) REFERENCES public.quote(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.order;
-- +goose StatementEnd
