-- name: MarkFeedFetched :exec
update feeds
set last_fetched_at = NOW(),
updated_at = NOW()
where id = $1;