-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_user_chapters (
    user_auth_id UUID NOT NULL,
    course_id UUID NOT NULL,
    chapter_id UUID NOT NULL,
    isFinished BOOLEAN default FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,


    CONSTRAINT fk_user_auth_id FOREIGN KEY (user_auth_id) REFERENCES t_users_auth(id) ON DELETE CASCADE,
    CONSTRAINT fk_course_id FOREIGN KEY (course_id) REFERENCES t_courses(id) ON DELETE CASCADE,
    CONSTRAINT fk_chapter_id FOREIGN KEY (chapter_id) REFERENCES t_chapters(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_user_chapters;
-- +goose StatementEnd
