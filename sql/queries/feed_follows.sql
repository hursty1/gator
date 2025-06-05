-- name: CreateFeedFollow :one
with inserted_feed_follows as (
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *)
select
ff.*,
f.name as feed_name,
u.name as user_name
from inserted_feed_follows ff
inner join users u on u.id = ff.user_id
inner join feeds f on f.id = ff.feed_id;

-- name: GetFeedFollowForUser :many
SELECT
ff.*,
u.name,
f.*
from feed_follows ff
inner join users u on u.id = ff.user_id
inner join feeds f on f.id = ff.feed_id
where ff.user_id = $1;


-- name: DeleteFeedFollows :exec
DELETE from feed_follows
where user_id = $1 and feed_id = $2;