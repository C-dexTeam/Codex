-- name: AddCourseToUser :exec
INSERT INTO t_user_courses 
    (user_auth_id, course_id, progress)
VALUES 
    (@user_auth_id, @course_id, 
     (
        SELECT COALESCE(
               (COUNT(CASE WHEN isFinished = TRUE THEN 1 END) * 100.0 / 
                NULLIF((SELECT COUNT(*) FROM t_chapters as c WHERE c.course_id = @course_id AND c.deleted_at IS NULL), 0)
               ), 0
           )
        FROM t_user_chapters as uc
        WHERE uc.user_auth_id = @user_auth_id AND uc.course_id = @course_id
     )
    );

-- name: UserCourses :many
SELECT 
    uc.user_auth_id,
    uc.course_id,
    c.title,
    uc.progress,
    COUNT(ucp.id) AS completed_chapters,
    (SELECT COUNT(*) FROM t_user_chapters WHERE course_id = uc.course_id) AS total_chapters
FROM 
    t_user_courses uc
INNER JOIN 
    t_courses c ON uc.course_id = c.id
LEFT JOIN 
    t_user_chapters ucp 
    ON ucp.course_id = uc.course_id AND ucp.isFinished = true
WHERE 
    uc.user_auth_id = @user_auth_id AND c.deleted_at IS NULL
GROUP BY 
    uc.user_auth_id, uc.course_id, c.title, c.description, uc.progress;

-- name: UserCourse :one
SELECT 
    user_auth_id, course_id, progress, created_at
FROM 
    t_user_courses
WHERE course_id = @course_id AND user_auth_id = @user_auth_id;

-- name: UpdateUserCourseProgress :one
UPDATE t_user_courses
SET progress = (
    SELECT COALESCE(
               (COUNT(CASE WHEN isFinished = TRUE THEN 1 END) * 100.0 / 
                (SELECT COUNT(*) FROM t_chapters as c WHERE c.course_id = @course_id AND c.deleted_at IS NULL)
               ), 0
           )
    FROM t_user_chapters as uc
    WHERE uc.user_auth_id = @user_auth_id AND uc.course_id = @course_id
)
WHERE 
    user_auth_id = @user_auth_id AND course_id = @course_id
RETURNING progress;

-- name: CheckUserCourseByID :one
SELECT 
CASE 
    WHEN EXISTS (
        SELECT 1 
        FROM t_user_courses AS l
        WHERE l.course_id = @course_id AND l.user_auth_id = @user_auth_id
    ) THEN true
    ELSE false
END AS exists;