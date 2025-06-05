-- name: CreatePosts :one
INSERT INTO posts (
    id,
    created_at,
    updated_at,
    title,
    url,
    description,
    feed_id
) values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
)
RETURNING *;



-- name: GetPostsForUser :many
with feeds as (
select
id,
feed_id
from feed_follows ff 
where ff.user_id = $1
)
select
*
from posts
where feed_id in (select feed_id from feeds)
order by created_at desc
limit $2;