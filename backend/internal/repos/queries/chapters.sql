-- name: GetChapters :many
SELECT 
    c.id, c.course_id, c.language_id, c.reward_id, c.reward_amount, c.title, c.description, c.content,
    c.func_name, c.frontend_template, c.docker_template, c.check_template, c.grants_experience, c.active,
    c.created_at, c.deleted_at
FROM 
    t_chapters as c
WHERE
    (sqlc.narg(id)::text IS NULL OR c.id = sqlc.narg(id)) AND
    (sqlc.narg(course_id)::text IS NULL OR c.course_id = sqlc.narg(course_id)) AND
    (sqlc.narg(language_id)::text IS NULL OR c.language_id = sqlc.narg(language_id)) AND
    (sqlc.narg(reward_id)::text IS NULL OR c.reward_id = sqlc.narg(reward_id)) AND
    (sqlc.narg(title)::text IS NULL OR title ILIKE '%' || sqlc.narg(title)::text || '%') AND
    deleted_at IS NULL
LIMIT @lim OFFSET @off;

-- name: GetChapterByID :one
SELECT
    c.id, 
    c.course_id, 
    c.language_id, 
    c.reward_id, 
    c.reward_amount, 
    c.title, 
    c.description, 
    c.content,
    c.func_name, 
    c.frontend_template, 
    c.docker_template, 
    c.check_template, 
    c.grants_experience, 
    c.active,
    c.created_at, 
    c.deleted_at,
    json_agg(
        json_build_object(
            'input_id', i.id,
            'input_value', i.value,
            'output_id', o.id,
            'output_value', o.value
        )
    ) AS tests
FROM 
    t_chapters as c
LEFT JOIN 
    t_tests as t ON t.chapter_id = c.id
LEFT JOIN 
    t_inputs as i ON i.test_id = t.id
LEFT JOIN 
    t_outputs as o ON o.test_id = t.id
WHERE
    c.id = @chapter_id
GROUP BY 
    c.id;


-- name: CreateChapter :one
INSERT INTO
    t_chapters (course_id, language_id, reward_id, reward_amount, title, description, content,
    func_name, frontend_template, docker_template, check_template, grants_experience, active)
VALUES
   (@course_id, @language_id, @reward_id, @reward_amount, @title, @description, @content,
    @func_name, @frontend_template, @docker_template, @check_template, @grants_experience, @active)
RETURNING id;


-- name: UpdateChapter :exec
UPDATE
    t_chapters
SET
    course_id = COALESCE(sqlc.narg(course_id), course_id),
    language_id = COALESCE(sqlc.narg(language_id), language_id),
    reward_id = COALESCE(sqlc.narg(reward_id), reward_id),
    reward_amount =  COALESCE(sqlc.narg(reward_amount), reward_amount),
    title =  COALESCE(sqlc.narg(title), title),
    description =  COALESCE(sqlc.narg(description), description),
    content =  COALESCE(sqlc.narg(content), content),
    func_name =  COALESCE(sqlc.narg(func_name), func_name),
    frontend_template =  COALESCE(sqlc.narg(frontend_template), frontend_template),
    docker_template =  COALESCE(sqlc.narg(docker_template), docker_template),
    check_template =  COALESCE(sqlc.narg(check_template), check_template),
    grants_experience =  COALESCE(sqlc.narg(grants_experience), grants_experience),
    active =  COALESCE(sqlc.narg(active), active)
WHERE
    id = @chapter_id;

-- name: SoftDeleteChapter :exec
UPDATE
    t_chapters
SET
    deleted_at = @deleted_at
WHERE  
    id = @chapter_id;
