-- name: GetAttributes :many
SELECT 
    a.id, a.reward_id, a.trait_type, a.value
FROM 
    t_attributes as a
WHERE
    (sqlc.narg(id)::text IS NULL OR us.id = sqlc.narg(id)) AND
    (sqlc.narg(reward_id)::text IS NULL OR us.reward_id = sqlc.narg(reward_id)) AND
    (sqlc.narg(trait_type)::text IS NULL OR trait_type ILIKE '%' || sqlc.narg(trait_type)::text || '%') AND
    (sqlc.narg(value)::text IS NULL OR value ILIKE '%' || sqlc.narg(value)::text || '%')
LIMIT @lim OFFSET @off;

-- name: GetAttributeByID :one
SELECT 
    a.id, a.reward_id, a.trait_type, a.value
FROM 
    t_attributes as a
WHERE
    id = @attribute_id;

-- name: UpdateAttribute :exec
UPDATE
    t_attributes
SET
    reward_id = COALESCE(sqlc.narg(reward_id)::TEXT, reward_id),
    trait_type = COALESCE(sqlc.narg(trait_type)::TEXT, trait_type),
    value = COALESCE(sqlc.narg(value)::TEXT, value)
WHERE
    id = @attribute_id;

-- name: DeleteAttribute :exec
DELETE FROM
    t_attributes
WHERE
    id = @attribute_id;