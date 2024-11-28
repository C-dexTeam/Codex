-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_courses (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    language_id UUID DEFAULT NULL,
    programming_language_id UUID DEFAULT NULL,
    reward_id UUID DEFAULT NULL,
    reward_amount INT NOT NULL,
    title VARCHAR(30),
    description TEXT,
    image_path varchar(60),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT fk_language_id FOREIGN KEY (language_id) REFERENCES t_languages(id) ON DELETE CASCADE,
    CONSTRAINT fk_reward_id FOREIGN KEY (reward_id) REFERENCES t_rewards(id) ON DELETE SET NULL,
    CONSTRAINT fk_pl_id FOREIGN KEY (programming_language_id) REFERENCES t_programming_language(id) ON DELETE SET NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_courses;
-- +goose StatementEnd
