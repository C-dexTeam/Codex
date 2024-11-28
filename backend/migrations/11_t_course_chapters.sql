-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_course_chapters (
    course_id UUID NOT NULL,
    chapter_id UUID NOT NULL,

    CONSTRAINT fk_course_id FOREIGN KEY (course_id) REFERENCES t_courses(id) ON DELETE CASCADE,
    CONSTRAINT fk_chapter_id FOREIGN KEY (chapter_id) REFERENCES t_chapters(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_course_chapters;
-- +goose StatementEnd
