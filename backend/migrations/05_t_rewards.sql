-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_rewards (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    reward_type VARCHAR(30) NOT NULL DEFAULT 'nft',
    symbol VARCHAR(30) NOT NULL,
    name VARCHAR(30) NOT NULL,
    description TEXT NOT NULL,
    image_path VARCHAR(60) NOT NULL,
    uri VARCHAR(120) NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_rewards;
-- +goose StatementEnd
