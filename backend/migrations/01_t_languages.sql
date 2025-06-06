-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_languages (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    value VARCHAR(30) NOT NULL,
    is_default BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO t_languages (value, is_default) VALUES ('EN', TRUE);
INSERT INTO t_languages (value, is_default) VALUES ('TR', FALSE);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_languages;
-- +goose StatementEnd
