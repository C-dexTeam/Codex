-- name: AddCourseToUser :exec
INSERT INTO t_user_courses 
    (user_auth_id, course_id, progress)
VALUES 
    (@user_auth_id, @course_id, 
     (
        SELECT 
            ROUND(100.0 * COUNT(*) / (SELECT COUNT(*) FROM t_user_chapters WHERE course_id = @course_id), 2)
        FROM t_user_chapters
        WHERE course_id = @course_id AND isFinished = true
     ));

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