-- name: GetChapters :many
SELECT 
    c.id, c.course_id, c.language_id, c.reward_id, c.title, c.description, c.content,
    c.func_name, c.frontend_template, c.docker_template,
    c.chapter_order, c.created_at, c.deleted_at
FROM 
    t_chapters as c
WHERE
    (sqlc.narg(id)::UUID IS NULL OR c.id = sqlc.narg(id)::UUID) AND
    (sqlc.narg(course_id)::UUID IS NULL OR c.course_id = sqlc.narg(course_id)::UUID) AND
    (sqlc.narg(language_id)::UUID IS NULL OR c.language_id = sqlc.narg(language_id)::UUID) AND
    (sqlc.narg(reward_id)::UUID IS NULL OR c.reward_id = sqlc.narg(reward_id)::UUID) AND
    (sqlc.narg(title)::text IS NULL OR c.title ILIKE '%' || sqlc.narg(title)::text || '%') AND
    deleted_at IS NULL
ORDER BY
    c.chapter_order  ASC
LIMIT 
    @lim OFFSET @off;

-- name: GetChapter :one
SELECT
    c.id, c.course_id, c.language_id, c.reward_id, c.title, c.description, c.content,
    c.func_name, c.frontend_template, c.docker_template,
    c.chapter_order, c.created_at, c.deleted_at
FROM
    t_chapters as c
WHERE
    c.id = @chapter_id;


-- name: CreateChapter :one
INSERT INTO
    t_chapters (course_id, language_id, reward_id, title, description, content,
    func_name, frontend_template, docker_template, chapter_order)
VALUES
   (@course_id, @language_id, @reward_id, @title, @description, @content,
    @func_name, @frontend_template, @docker_template, @chapter_order)
RETURNING id;


-- name: UpdateChapter :exec
UPDATE
    t_chapters
SET
    course_id = COALESCE(sqlc.narg(course_id), course_id),
    language_id = COALESCE(sqlc.narg(language_id), language_id),
    reward_id = COALESCE(sqlc.narg(reward_id), reward_id),
    title =  COALESCE(sqlc.narg(title), title),
    description =  COALESCE(sqlc.narg(description), description),
    content =  COALESCE(sqlc.narg(content), content),
    func_name =  COALESCE(sqlc.narg(func_name), func_name),
    frontend_template =  COALESCE(sqlc.narg(frontend_template), frontend_template),
    docker_template =  COALESCE(sqlc.narg(docker_template), docker_template)
WHERE
    id = @chapter_id;

-- name: SoftDeleteChapter :exec
UPDATE
    t_chapters
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE  
    id = @chapter_id;

-- name: CheckChapterByID :one
SELECT 
CASE 
    WHEN EXISTS (
        SELECT 1 
        FROM t_chapters AS l
        WHERE l.id = @chapter_id AND l.deleted_at IS NULL
    ) THEN true
    ELSE false
END AS exists;