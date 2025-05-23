-- name: GetCourses :many
SELECT 
    id, language_id, programming_language_id, reward_id, title, c.description, c.image_path, 
    (SELECT COUNT(*) FROM t_chapters as ch WHERE ch.course_id = c.id) as chapter_count,
    c.created_at, c.deleted_at
FROM 
    t_courses as c
WHERE
    (sqlc.narg(id)::UUID IS NULL OR c.id = sqlc.narg(id)::UUID) AND
    (sqlc.narg(language_id)::UUID IS NULL OR c.language_id = sqlc.narg(language_id)::UUID) AND
    (sqlc.narg(programming_language_id)::UUID IS NULL OR c.programming_language_id = sqlc.narg(programming_language_id)::UUID) AND
    (sqlc.narg(reward_id)::UUID IS NULL OR c.reward_id = sqlc.narg(reward_id):: UUID) AND
    (sqlc.narg(title)::text IS NULL OR c.title ILIKE '%' || sqlc.narg(title)::text || '%') AND
    deleted_at IS NULL
LIMIT 
    @lim OFFSET @off;

-- name: GetCourse :one
SELECT 
    id, language_id, programming_language_id, reward_id, title, c.description, c.image_path, 
    (SELECT COUNT(*) FROM t_chapters as ch WHERE ch.course_id = c.id) as chapter_count,
    c.created_at, c.deleted_at
FROM 
    t_courses as c
WHERE
    c.id = @course_id;

-- name: GetTopCourses :many
SELECT 
    c.id,
    c.language_id,
    c.programming_language_id,
    c.reward_id,
    c.title,
    c.description,
    c.image_path,
    (SELECT COUNT(*) FROM t_chapters AS ch WHERE ch.course_id = c.id) AS chapter_count,
    c.created_at,
    c.deleted_at
FROM 
    t_user_courses AS uc
JOIN 
    t_courses AS c
ON 
    uc.course_id = c.id
GROUP BY 
    c.id, c.language_id, c.programming_language_id, c.reward_id, c.title, 
    c.description, c.image_path, c.created_at, c.deleted_at
ORDER BY 
    COUNT(uc.user_auth_id) DESC
LIMIT @lim OFFSET @off;


-- name: CreateCourse :one
INSERT INTO
    t_courses (language_id, programming_language_id, reward_id, title, description, image_path)
VALUES
    (@language_id, @programming_language_id, @reward_id, @title, @description, @image_path)
RETURNING id;

-- name: UpdateCourse :exec
UPDATE
    t_courses
SET
    language_id =  COALESCE(sqlc.narg(language_id)::UUID, language_id),
    programming_language_id =  COALESCE(sqlc.narg(programming_language_id)::UUID, programming_language_id),
    reward_id =  COALESCE(sqlc.narg(reward_id)::UUID, reward_id),
    title =  COALESCE(sqlc.narg(title)::TEXT, title),
    description =  COALESCE(sqlc.narg(description)::TEXT, description),
    image_path =  COALESCE(sqlc.narg(image_path)::TEXT, image_path)
WHERE
    id = @course_id;

-- name: SoftDeleteCourse :exec
UPDATE
    t_courses
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE  
    id = @course_id;

-- name: DeleteCourse :exec
DELETE FROM
    t_courses
WHERE
    id = @course_id;

-- name: CheckCourseByID :one
SELECT 
CASE 
    WHEN EXISTS (
        SELECT 1 
        FROM t_courses AS l
        WHERE l.id = @course_id AND l.deleted_at IS NULL
    ) THEN true
    ELSE false
END AS exists;

-- name: CourseCount :one
SELECT COUNT(*) FROM t_courses WHERE deleted_at IS NULL;