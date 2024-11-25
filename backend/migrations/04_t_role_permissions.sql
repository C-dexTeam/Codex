-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS t_role_permissions (
    role_id UUID NOT NULL,
    permission_id UUID NOT NULL,
    
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES t_roles(id) ON DELETE CASCADE,
    CONSTRAINT fk_permission FOREIGN KEY (permission_id) REFERENCES t_permissions(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_role_permissions;
-- +goose StatementEnd
