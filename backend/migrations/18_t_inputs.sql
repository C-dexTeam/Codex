-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_inputs (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    test_id UUID NOT NULL,
    
    CONSTRAINT fk_test_id FOREIGN KEY (test_id) REFERENCES t_tests(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_inputs;
-- +goose StatementEnd
