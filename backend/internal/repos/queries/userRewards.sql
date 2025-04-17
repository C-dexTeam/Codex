-- name: AddRewardToUser :exec
INSERT INTO t_user_rewards
    (user_auth_id, course_id, chapter_id, reward_id)
VALUES
    (@user_auth_id, @course_id, @chapter_id, @reward_id);

-- name: CheckUserReward :one
SELECT 
CASE 
    WHEN EXISTS (
        SELECT 1 
        FROM t_user_rewards AS l
        WHERE
            l.user_auth_id = @user_auth_id AND course_id = @course_id AND chapter_id = @chapter_id AND reward_id = @reward_id
    ) THEN true
    ELSE false
END AS exists;

-- name: UserRewards :many
SELECT
    r.id, r.symbol, r.name, r.description, r.image_path, r.uri, 
    ur.created_at AS earned_date
FROM
    t_rewards AS r
INNER JOIN 
    t_user_rewards AS ur ON r.id = ur.reward_id
WHERE 
    ur.user_auth_id = @user_auth_id
ORDER BY 
    ur.created_at DESC
LIMIT 
    @lim OFFSET @off;
