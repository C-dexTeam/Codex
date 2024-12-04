-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_attributes (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    reward_id UUID DEFAULT NULL,
    trait_type VARCHAR(30) NOT NULL,
    value VARCHAR(30) NOT NULL,

    CONSTRAINT fk_reward_id FOREIGN KEY (reward_id) REFERENCES t_rewards(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_attributes;
-- +goose StatementEnd
