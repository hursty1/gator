-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: FetchAllFeeds :many
select
f.*,
u.name as user_name
from feeds f
inner join users u on u.id = f.user_id;

-- name: FetchFeedByUrl :one
select
f.*,
u.name as user_name
from feeds f
inner join users u on u.id = f.user_id
where f.url = $1;