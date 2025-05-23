-- name: GetRewards :many
SELECT 
    r.id, r.symbol, r.name, r.description, r.seller_fee, r.image_path, r.uri 
FROM 
    t_rewards as r
WHERE
    (sqlc.narg(id)::UUID IS NULL OR r.id = sqlc.narg(id)::UUID) AND
    (sqlc.narg(symbol)::text IS NULL OR r.symbol ILIKE '%' || sqlc.narg(symbol)::text || '%') AND
    (sqlc.narg(name)::text IS NULL OR r.name ILIKE '%' || sqlc.narg(name)::text || '%') AND
    (sqlc.narg(description)::text IS NULL OR r.description ILIKE '%' || sqlc.narg(description)::text || '%')
LIMIT @lim OFFSET @off;

-- name: GetReward :one
SELECT 
    r.id, r.symbol, r.name, r.description, r.seller_fee, r.image_path, r.uri
FROM 
    t_rewards as r
WHERE
    r.id = @reward_id;

-- name: CreateReward :one
INSERT INTO
    t_rewards (symbol, name, description, image_path, uri, seller_fee)
VALUES
    (@symbol, @name, @description, @image_path, @uri, @seller_fee)
RETURNING id;

-- name: UpdateReward :exec
UPDATE
    t_rewards
SET
    symbol =  COALESCE(sqlc.narg(symbol)::TEXT, symbol),
    name =  COALESCE(sqlc.narg(name)::TEXT, name),
    description =  COALESCE(sqlc.narg(description)::TEXT, description),
    image_path =  COALESCE(sqlc.narg(image_path)::TEXT, image_path),
    uri =  COALESCE(sqlc.narg(uri)::TEXT, uri),
    seller_fee =  COALESCE(sqlc.narg(seller_fee)::INTEGER, seller_fee)
WHERE
    id = @reward_id;

-- name: DeleteReward :exec
DELETE FROM
    t_rewards
WHERE
    id = @reward_id;

-- name: CheckRewardByID :one
SELECT 
CASE 
    WHEN EXISTS (
        SELECT 1 
        FROM t_rewards AS l
        WHERE l.id = @reward_id 
    ) THEN true
    ELSE false
END AS exists;

-- name: RewardCount :one
SELECT COUNT(*) FROM t_rewards;