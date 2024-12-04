-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_outputs (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    input_id UUID NOT NULL,
    value TEXT,
    
    CONSTRAINT fk_test_id FOREIGN KEY (input_id) REFERENCES t_inputs(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_outputs;
-- +goose StatementEnd
