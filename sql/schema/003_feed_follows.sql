-- +goose up
CREATE TABLE feed_follows (
    id uuid PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    
    user_id UUID NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,
    
    feed_id UUID NOT NULL,
    CONSTRAINT fk_feed
        FOREIGN KEY (feed_id)
        REFERENCES feeds(id)
        ON DELETE CASCADE,
    
    CONSTRAINT unique_user_feed UNIQUE (user_id, feed_id)
);

-- +goose down
DROP TABLE feed_follows;