-- +goose Up
CREATE TABLE quotes (
    id VARCHAR(36) PRIMARY KEY,
    author VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS quotes; 