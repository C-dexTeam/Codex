-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_programming_languages (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    language_id UUID DEFAULT NULL,
    name VARCHAR(30) NOT NULL,
    description TEXT NOT NULL,
    download_cmd VARCHAR(256) NOT NULL,
    compile_cmd VARCHAR(256) NOT NULL,
    image_path VARCHAR(60) NOT NULL,
    file_extention VARCHAR(10) NOT NULL,
    monaco_editor VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_language_id FOREIGN KEY (language_id) REFERENCES t_languages(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_programming_languages;
-- +goose StatementEnd
