-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_reward_attributes (
    reward_id UUID NOT NULL,
    attribute_id UUID NOT NULL,

    CONSTRAINT fk_reward_id FOREIGN KEY (reward_id) REFERENCES t_rewards(id) ON DELETE CASCADE,
    CONSTRAINT fk_attr_id FOREIGN KEY (attribute_id) REFERENCES t_attributes(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_reward_attributes;
-- +goose StatementEnd
