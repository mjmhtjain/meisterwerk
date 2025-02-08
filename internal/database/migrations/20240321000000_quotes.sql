-- +goose Up
CREATE TABLE quote (
    id VARCHAR(255) PRIMARY KEY,
    author VARCHAR(255) NOT NULL,
    customer_name VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS quote; 