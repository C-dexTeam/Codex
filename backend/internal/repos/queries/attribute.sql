-- name: GetAttributes :many
SELECT 
    a.id, a.reward_id, a.trait_type, a.value
FROM 
    t_attributes as a
WHERE
    (sqlc.narg(id)::UUID IS NULL OR a.id = sqlc.narg(id)::UUID) AND
    (sqlc.narg(reward_id)::UUID IS NULL OR a.reward_id = sqlc.narg(reward_id)::UUID) AND
    (sqlc.narg(trait_type)::text IS NULL OR a.trait_type ILIKE '%' || sqlc.narg(trait_type)::text || '%') AND
    (sqlc.narg(value)::text IS NULL OR a.value ILIKE '%' || sqlc.narg(value)::text || '%')
LIMIT @lim OFFSET @off;

-- name: GetAttributeByID :one
SELECT 
    a.id, a.reward_id, a.trait_type, a.value
FROM 
    t_attributes as a
WHERE
    id = @attribute_id;

-- name: CreateAttribute :one
INSERT INTO 
    t_attributes (reward_id, trait_type, value)
VALUES
    (@reward_id, @trait_type, @value)
RETURNING id;

-- name: UpdateAttribute :exec
UPDATE
    t_attributes
SET
    reward_id = COALESCE(sqlc.narg(reward_id)::UUID, reward_id),
    trait_type = COALESCE(sqlc.narg(trait_type)::TEXT, trait_type),
    value = COALESCE(sqlc.narg(value)::TEXT, value)
WHERE
    id = @attribute_id;

-- name: DeleteAttribute :exec
DELETE FROM
    t_attributes
WHERE
    id = @attribute_id;

-- name: AttributeCount :one
SELECT COUNT(*) FROM t_attributes;