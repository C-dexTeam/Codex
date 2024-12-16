-- name: GetUsersProfile :many
SELECT up.id, up.user_auth_id, up.role_id, up.name, up.surname, up.level, up.experience, up.next_level_exp,
       up.created_at, up.deleted_at 
FROM t_users_profile as up
WHERE
    (sqlc.narg(id)::UUID IS NULL OR up.id = sqlc.narg(id)::UUID) AND
    (sqlc.narg(user_auth_id)::UUID IS NULL OR up.user_auth_id = sqlc.narg(user_auth_id)::UUID) AND
    (sqlc.narg(role_id)::UUID IS NULL OR up.role_id = sqlc.narg(role_id)::UUID) AND
    (sqlc.narg(name)::TEXT IS NULL OR up.name ILIKE '%' || sqlc.narg(name)::TEXT || '%') AND
    (sqlc.narg(surname)::TEXT IS NULL OR up.surname ILIKE '%' || sqlc.narg(surname)::TEXT || '%') AND
    (sqlc.narg(level)::INTEGER IS NULL OR up.level = sqlc.narg(level)::INTEGER) AND
    (sqlc.narg(experience)::INTEGER IS NULL OR up.experience = sqlc.narg(experience)::INTEGER) AND
    (sqlc.narg(next_level_exp)::INTEGER IS NULL OR up.next_level_exp = sqlc.narg(next_level_exp)::INTEGER)
LIMIT @lim OFFSET @off;

-- name: GetUserProfileByID :one
SELECT 
    up.id, up.user_auth_id, up.role_id, up.name, up.surname, up.level, up.experience, up.next_level_exp, up.created_at, up.deleted_at 
FROM t_users_profile as up
WHERE
    up.id = @id;

-- name: CreateUserProfile :exec
INSERT INTO
    t_users_profile (user_auth_id, role_id, name, surname)
VALUES
    (@user_auth_id, @role_id, @name, @surname);

-- name: ChangeUserRole :exec
UPDATE
    t_users_profile
SET
    role_id = @role_id
WHERE
    id = @user_profile_id;

-- name: ChangeUserLevel :exec
UPDATE
    t_users_profile
SET
    level =  COALESCE(sqlc.narg(level)::INTEGER, level),
    experience =  COALESCE(sqlc.narg(experience)::INTEGER, experience),
    next_level_Exp =  COALESCE(sqlc.narg(next_level_exp)::INTEGER, next_level_Exp)
WHERE
    id = @user_profile_id;

-- name: SoftDeleteUserProfile :exec
UPDATE
    t_users_profile
SET
    deleted_at = @deleted_at
WHERE
    id = @user_profile_id;

-- name: UpdateUserProfile :exec
UPDATE
    t_users_profile
SET
    name = COALESCE(sqlc.narg(name)::TEXT, name),
    surname =  COALESCE(sqlc.narg(surname)::TEXT, surname)
WHERE
    id = @user_profile_id;
