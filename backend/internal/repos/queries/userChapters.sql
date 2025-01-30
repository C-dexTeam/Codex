-- name: AddChapterToUser :exec
INSERT INTO t_user_chapters 
    (user_auth_id, course_id, chapter_id, isFinished)
VALUES 
    (@user_auth_id, @course_id, @chapter_id, FALSE);