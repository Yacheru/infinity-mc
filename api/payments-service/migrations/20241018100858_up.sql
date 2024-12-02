-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    email VARCHAR(50) NOT NULL,
    nickname VARCHAR(50) NOT NULL,
    service VARCHAR(50) NOT NULL,
    price VARCHAR(5) NOT NULL,
    duration BIGINT NOT NULL,
    created_at timestamptz default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments;
-- +goose StatementEnd
