-- name: GetUserAuths :many
SELECT 
    us.id, us.public_key, us.username, us.email, us.deleted_at
FROM 
    t_users_auth as us
WHERE
    (sqlc.narg(id)::text IS NULL OR us.id = sqlc.narg(id)) AND
    (sqlc.narg(public_key)::text IS NULL OR us.public_key = sqlc.narg(public_key)) AND
    (sqlc.narg(username)::text IS NULL OR username ILIKE '%' || sqlc.narg(username)::text || '%') AND
    (sqlc.narg(email)::text IS NULL OR email ILIKE '%' || sqlc.narg(email)::text || '%') AND
    deleted_at IS NULL
LIMIT @lim OFFSET @off;

-- name: GetUserAuthByID :one
SELECT 
    id, public_key, username, email 
FROM 
    t_users_auth
WHERE 
    id = @user_auth_id;

-- name: CreateUserAuth :exec
INSERT INTO t_users_auth 
    (public_key, username, email, password)
VALUES 
    (@public_key, @username, @email, @password);

-- name: UpdateUserAuth :exec
UPDATE
    t_users_auth
SET
    public_key = COALESCE(sqlc.narg(public_key)::TEXT, public_key),
    username = COALESCE(sqlc.narg(username)::TEXT, username),
    email = COALESCE(sqlc.narg(email)::TEXT, email)
WHERE
    id = @user_auth_id;

-- name: SoftDeleteUser :exec
UPDATE
    t_users_auth
SET
    deleted_at = @deleted_at
WHERE
    id = @user_auth_id;
