-- +goose up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT not null,
    url TEXT not null,
    description TEXT not null,
    feed_id UUID NOT NULL,
    CONSTRAINT fk_feed
        FOREIGN KEY (feed_id)
        REFERENCES feeds(id)
        ON DELETE CASCADE
);

-- +goose down
DROP TABLE posts;