-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_user_profiles (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    role_id UUID NOT NULL,
    name varchar(30),
    surname varchar(30),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES t_users(id) ON DELETE CASCADE,
    CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES t_roles(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_users;
-- +goose StatementEnd
