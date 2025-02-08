-- +goose Up
-- +goose StatementBegin
INSERT INTO product (id, name, price, tax) VALUES ('1', 'Product 1', 100.00, 10.00);
INSERT INTO product (id, name, price, tax) VALUES ('2', 'Product 2', 200.00, 15.50);
INSERT INTO product (id, name, price, tax) VALUES ('3', 'Product 3', 300.00, 12.25);
INSERT INTO product (id, name, price, tax) VALUES ('4', 'Product 4', 150.00, 8.00);
INSERT INTO product (id, name, price, tax) VALUES ('5', 'Product 5', 250.00, 20.00);
INSERT INTO product (id, name, price, tax) VALUES ('6', 'Product 6', 350.00, 18.00);
INSERT INTO product (id, name, price, tax) VALUES ('7', 'Product 7', 450.00, 22.50);
INSERT INTO product (id, name, price, tax) VALUES ('8', 'Product 8', 550.00, 25.00);
INSERT INTO product (id, name, price, tax) VALUES ('9', 'Product 9', 650.00, 30.00);
INSERT INTO product (id, name, price, tax) VALUES ('10', 'Product 10', 750.00, 35.00);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM product WHERE id IN ('1', '2', '3', '4', '5', '6', '7', '8', '9', '10');
-- +goose StatementEnd
