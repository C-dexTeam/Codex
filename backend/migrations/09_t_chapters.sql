-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_chapters (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    course_id UUID NOT NULL,
    language_id UUID NOT NULL,
    reward_id UUID DEFAULT NULL,
    title VARCHAR(30) NOT NULL,
    description TEXT NOT NULL,
    content TEXT NOT NULL,
    func_name VARCHAR(30) NOT NULL,
    frontend_template TEXT NOT NULL,
    docker_template TEXT NOT NULL,
    chapter_order INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,

    CONSTRAINT fk_reward_id FOREIGN KEY (reward_id) REFERENCES t_rewards(id) ON DELETE SET NULL,
    CONSTRAINT fk_language_id FOREIGN KEY (language_id) REFERENCES t_languages(id) ON DELETE CASCADE,
    CONSTRAINT fk_course_id FOREIGN KEY (course_id) REFERENCES t_courses(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_chapters;
-- +goose StatementEnd
