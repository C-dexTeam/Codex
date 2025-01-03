-- name: GetLanguages :many
SELECT
    l.id, l.value, l.is_default
FROM 
    t_languages as l
WHERE
    (sqlc.narg(id)::UUID IS NULL OR l.id = sqlc.narg(id)::UUID) AND
    (sqlc.narg(value)::text IS NULL OR l.value = sqlc.narg(value)::TEXT)
LIMIT @lim OFFSET @off;

-- name: GetLanguageByID :one
SELECT
    l.id, l.value, l.is_default
FROM 
    t_languages as l
WHERE
    l.id = @language_id;

-- name: GetDefaultLanguage :one
SELECT
    id, value, is_default
FROM
    t_languages
WHERE 
    is_default = True;

-- name: GetLanguageByValue :one
SELECT
    l.id, l.value, l.is_default
FROM 
    t_languages as l
WHERE
    l.value = @language_value;

-- name: CheckLanguageByID :one
SELECT 
CASE 
    WHEN EXISTS (
        SELECT 1 
        FROM t_languages AS l
        WHERE l.id = @language_id
    ) THEN true
    ELSE false
END AS exists;
