-- name: GetCourses :many
SELECT 
    c.id, c.language_id, c.programming_language_id, c.reward_id, c.reward_amount, c.title,
    c.description, c.image_path, c.created_at, c.deleted_at
FROM 
    t_courses as c
WHERE
    (sqlc.narg(id)::text IS NULL OR c.id = sqlc.narg(id)) AND
    (sqlc.narg(language_id)::text IS NULL OR c.language_id = sqlc.narg(language_id)) AND
    (sqlc.narg(programming_language_id)::text IS NULL OR c.programming_language_id = sqlc.narg(programming_language_id)) AND
    (sqlc.narg(reward_id)::text IS NULL OR c.reward_id = sqlc.narg(reward_id)) AND
    (sqlc.narg(title)::text IS NULL OR title ILIKE '%' || sqlc.narg(title)::text || '%') AND
    (sqlc.narg(description)::text IS NULL OR description ILIKE '%' || sqlc.narg(description)::text || '%') AND
    deleted_at IS NULL
LIMIT @lim OFFSET @off;

-- name: GetCourseByID :many
SELECT 
    c.id, 
    c.language_id, 
    c.programming_language_id, 
    c.reward_id, 
    c.reward_amount, 
    c.title,
    c.description, 
    c.created_at, 
    c.deleted_at,
    (
        SELECT 
            json_agg(
                json_build_object(
                    'id', ch.id,
                    'language_id', ch.language_id,
                    'reward_id', ch.reward_id,
                    'reward_amount', ch.reward_amount,
                    'title', ch.title,
                    'description', ch.description,
                    'content', ch.content,
                    'func_name', ch.func_name,
                    'frontend_template', ch.frontend_template,
                    'docker_template', ch.docker_template,
                    'check_template', ch.check_template,
                    'grants_experience', ch.grants_experience,
                    'active', ch.active,
                    'created_at', ch.created_at,
                    'deleted_at', ch.deleted_at
                )
            )
        FROM 
            t_chapters as ch
        WHERE 
            ch.course_id = c.id
        LIMIT @lim OFFSET @off
    ) AS chapters
FROM 
    t_courses as c
WHERE
    c.id = @course_id;

-- name: CreateCourse :exec
INSERT INTO
    t_courses (language_id, programming_language_id, reward_id, reward_amount, title, description, image_path)
VALUES
    (@language_id, @programming_language_id, @reward_id, @reward_amount, @title, @description, @image_path);

-- name: UpdateCourse :exec
UPDATE
    t_courses
SET
    language_id =  COALESCE(sqlc.narg(language_id)::TEXT, language_id),
    programming_language_id =  COALESCE(sqlc.narg(programming_language_id)::TEXT, programming_language_id),
    reward_id =  COALESCE(sqlc.narg(reward_id)::TEXT, reward_id),
    reward_amount =  COALESCE(sqlc.narg(reward_amount)::INTEGER, reward_amount),
    title =  COALESCE(sqlc.narg(title)::TEXT, title),
    description =  COALESCE(sqlc.narg(description)::TEXT, description),
    image_path =  COALESCE(sqlc.narg(image_path)::TEXT, image_path)
WHERE
    id = @course_id;

-- name: SoftDeleteCourse :exec
UPDATE
    t_courses
SET
    deleted_at = @deleted_at
WHERE  
    id = @course_id;
