-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_users_auth (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    public_key VARCHAR(128),
    username VARCHAR(30),
    email VARCHAR(40),
    password VARCHAR(255),
    deleted_at TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_users_auth;
-- +goose StatementEnd
