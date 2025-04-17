-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_rewards (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    symbol VARCHAR(30) NOT NULL,
    name VARCHAR(30) NOT NULL,
    description TEXT NOT NULL,
    seller_fee INT NOT NULL,
    image_path VARCHAR(60) DEFAULT NULL,
    uri VARCHAR(120) DEFAULT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_rewards;
-- +goose StatementEnd
