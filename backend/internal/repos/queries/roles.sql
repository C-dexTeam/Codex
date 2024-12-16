-- name: GetRoles :many
SELECT
    r.id, r.name
FROM 
    t_roles as r
WHERE
    (sqlc.narg(id)::text IS NULL OR us.id = sqlc.narg(id)) AND
    (sqlc.narg(name)::text IS NULL OR name ILIKE '%' || sqlc.narg(name)::text || '%')
LIMIT @lim OFFSET @off;

-- name: GetRoleByName :one
SELECT
    r.id, r.name
FROM 
    t_roles as r
WHERE
    r.name = @role_name;

-- name: GetRoleByID :one
SELECT
    r.id, r.name
FROM 
    t_roles as r
WHERE
    r.id = @role_id;