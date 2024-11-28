-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_chapters (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    language_id UUID DEFAULT NULL,
    title VARCHAR(30),
    description TEXT NOT NULL,
    content TEXT NOT NULL,
    func_name VARCHAR(30) NOT NULL,
    frontend_template TEXT NOT NULL,
    docker_template TEXT NOT NULL,
    check_template TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    CONSTRAINT fk_language_id FOREIGN KEY (language_id) REFERENCES t_languages(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_chapters;
-- +goose StatementEnd
