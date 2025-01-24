-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_users_profile (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_auth_id UUID NOT NULL,
    role_id UUID NOT NULL,
    name varchar(30),
    surname varchar(30),
    level INT DEFAULT 1,
    experience INT DEFAULT 0,
    next_level_exp INT DEFAULT 100,
    streak INT DEFAULT 0,
    last_streak_date TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT fk_user_auth_id FOREIGN KEY (user_auth_id) REFERENCES t_users_auth(id) ON DELETE CASCADE,
    CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES t_roles(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_users_profile;
-- +goose StatementEnd
