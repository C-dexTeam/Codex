-- name: AddChapterToUser :exec
INSERT INTO t_user_chapters 
    (user_auth_id, course_id, chapter_id, isFinished)
VALUES 
    (@user_auth_id, @course_id, @chapter_id, FALSE);

-- name: UpdateUserChapter :exec
UPDATE
    t_user_chapters
SET
    isFinished = TRUE
WHERE
    user_auth_id = @user_auth_id AND course_id = @course_id AND chapter_id =  @chapter_id;