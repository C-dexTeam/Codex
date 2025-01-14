-- name: GetTests :many
SELECT
    t.id, t.chapter_id, t.input_value, t.output_value
FROM 
    t_tests AS t
WHERE
    (sqlc.narg(id)::UUID IS NULL OR t.id = sqlc.narg(id)::UUID) AND
    (sqlc.narg(chapter_id)::UUID IS NULL OR t.chapter_id = sqlc.narg(chapter_id)::UUID)
LIMIT @lim OFFSET @off;

-- name: GetTest :one
SELECT
    t.id, t.chapter_id, t.input_value, t.output_value
FROM 
    t_tests as t
WHERE
    t.id = @test_id;

-- name: CreateTest :one
INSERT INTO t_tests (chapter_id, input_value, output_value)
VALUES (@chapter_id, @input_value, @output_value)
RETURNING id;

-- name: DeleteTest :exec
DELETE FROM 
    t_tests
WHERE 
    id = @test_id;

-- name: UpdateTest :exec
UPDATE
    t_tests
SET
    input_value =  COALESCE(sqlc.narg(input_value)::TEXT, input_value),
    output_value =  COALESCE(sqlc.narg(output_value)::TEXT, output_value)
WHERE
    id = @test_id;

-- name: CheckTestByID :one
SELECT 
CASE 
    WHEN EXISTS (
        SELECT 1 
        FROM t_tests AS l
        WHERE l.id = @test_id 
    ) THEN true
    ELSE false
END AS exists;

