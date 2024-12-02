-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    uuid UUID NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    nickname VARCHAR(50) NOT NULL UNIQUE,
    ip CIDR NOT NULL,
    password VARCHAR NOT NULL,
    refresh_token VARCHAR DEFAULT NULL,
    expires_in BIGINT DEFAULT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
