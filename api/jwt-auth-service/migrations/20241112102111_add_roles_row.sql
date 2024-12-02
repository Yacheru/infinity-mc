-- +goose Up
-- +goose StatementBegin
CREATE TYPE roles AS ENUM ('player', 'admin');

ALTER TABLE IF EXISTS users ADD COLUMN IF NOT EXISTS role roles DEFAULT 'player';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS roles;
-- +goose StatementEnd
