-- name: GetRewards :many
SELECT 
    r.id, r.reward_type, r.symbol, r.name, r.description, r.image_path, r.uri
FROM 
    t_rewards as r
WHERE
    (sqlc.narg(id)::text IS NULL OR us.id = sqlc.narg(id)) AND
    (sqlc.narg(reward_type)::text IS NULL OR us.reward_type = sqlc.narg(reward_type)) AND
    (sqlc.narg(symbol)::text IS NULL OR symbol ILIKE '%' || sqlc.narg(symbol)::text || '%') AND
    (sqlc.narg(name)::text IS NULL OR name ILIKE '%' || sqlc.narg(name)::text || '%') AND
    (sqlc.narg(description)::text IS NULL OR description ILIKE '%' || sqlc.narg(description)::text || '%')
LIMIT @lim OFFSET @off;

-- name: GetReward :one
SELECT 
    r.id, r.reward_type, r.symbol, r.name, r.description, r.image_path, r.uri,
    (SELECT * FROM t_attributes LIMIT @lim OFFSET @off) AS attributes
FROM 
    t_rewards as r
WHERE
    r.id = @reward_id;

-- name: CreateReward :exec
INSERT INTO
    t_rewards (reward_type, symbol, name, description, image_path, uri)
VALUES
    (@reward_type, @symbol, @name, @description, @image_path, @uri);

-- name: UpdateReward :exec
UPDATE
    t_rewards
SET
    reward_type =  COALESCE(sqlc.narg(reward_type)::TEXT, reward_type),
    symbol =  COALESCE(sqlc.narg(symbol)::TEXT, symbol),
    name =  COALESCE(sqlc.narg(name)::TEXT, name),
    description =  COALESCE(sqlc.narg(description)::TEXT, description),
    image_path =  COALESCE(sqlc.narg(image_path)::TEXT, image_path),
    uri =  COALESCE(sqlc.narg(uri)::TEXT, uri)
WHERE
    id = @reward_id;

-- name: DeleteReward :exec
DELETE FROM
    t_rewards
WHERE
    id = @reward_id;