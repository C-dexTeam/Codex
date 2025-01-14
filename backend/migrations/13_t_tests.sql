-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_tests (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    chapter_id UUID NOT NULL,
    input_value TEXT NOT NULL,
    output_value TEXT NOT NULL,
    
    CONSTRAINT fk_chapter_id FOREIGN KEY (chapter_id) REFERENCES t_chapters(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_tests;
-- +goose StatementEnd
