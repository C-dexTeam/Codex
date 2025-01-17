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
