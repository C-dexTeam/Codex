-- name: GetTests :many
SELECT
    t.id, i.value, o.value
FROM 
    t_tests AS t
INNER JOIN
    t_inputs i
ON i.test_id = t.id
INNER JOIN 
    t_outputs o
ON o.test_id = t.id
WHERE
    (sqlc.narg(id)::UUID IS NULL OR t.id = sqlc.narg(id)::UUID)
LIMIT @lim OFFSET @off;

-- name: GetTestByID :one
SELECT
    t.id, i.value, o.value
FROM 
    t_tests AS t
INNER JOIN
    t_inputs i
ON i.test_id = t.id
INNER JOIN 
    t_outputs o
ON o.test_id = t.id
WHERE
    t.id = @test_id;

-- name: CreateTest :exec
BEGIN;

-- 1. Adım: Test ekleme
INSERT INTO t_tests (chapter_id)
VALUES (@chapter_id)
RETURNING id;

-- 2. Adım: Input ekleme
INSERT INTO t_inputs (test_id, value)
VALUES ((SELECT id FROM t_tests WHERE chapter_id = @chapter_id), @input_value);

-- 3. Adım: Output ekleme
INSERT INTO t_outputs (test_id, value)
VALUES ((SELECT id FROM t_tests WHERE chapter_id = @chapter_id), @output_value);

-- name: DeleteTest :exec
DELETE FROM 
    t_tests
WHERE 
    id = @test_id;
