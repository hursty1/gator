-- name: GetNextFeedToFetch :one
select
*
from feeds
order by last_fetched_at desc 
limit 1;