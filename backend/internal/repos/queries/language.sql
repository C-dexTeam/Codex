-- name: GetLanguages :many
SELECT
    l.id, l.value
FROM t_languages as l
WHERE
    (sqlc.narg(id)::text IS NULL OR us.id = sqlc.narg(id)) AND
    (sqlc.narg(value)::text IS NULL OR value ILIKE '%' || sqlc.narg(value)::text || '%')
LIMIT @lim OFFSET @off;