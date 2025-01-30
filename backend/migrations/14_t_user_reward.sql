-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_user_rewards (
    user_auth_id UUID NOT NULL,
    course_id UUID NOT NULL,
    chapter_id UUID NOT NULL,
    reward_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,


    CONSTRAINT fk_user_auth_id FOREIGN KEY (user_auth_id) REFERENCES t_users_auth(id) ON DELETE CASCADE,
    CONSTRAINT fk_course_id FOREIGN KEY (course_id) REFERENCES t_courses(id) ON DELETE CASCADE,
    CONSTRAINT fk_chapter_id FOREIGN KEY (chapter_id) REFERENCES t_chapters(id) ON DELETE CASCADE,
    CONSTRAINT fk_reward_id FOREIGN KEY (reward_id) REFERENCES t_rewards(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_reward UNIQUE (user_auth_id, reward_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_user_rewards;
-- +goose StatementEnd
