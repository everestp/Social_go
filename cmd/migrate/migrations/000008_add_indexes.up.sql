-- Enable extension (safe to run always)
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Comments and titles are normal text → pg_trgm works fine
CREATE INDEX IF NOT EXISTS idx_comments_content 
    ON comments USING gin (content gin_trgm_ops);

CREATE INDEX IF NOT EXISTS idx_post_title 
    ON posts USING gin (title gin_trgm_ops);

-- tags is text[] / varchar[] → gin_trgm_ops does NOT work on arrays
-- Use the built-in array GIN index instead (perfect for "WHERE tags @> ARRAY['golang']'")
DROP INDEX IF EXISTS idx_post_tags;  -- remove the broken one first if it exists
CREATE INDEX IF NOT EXISTS idx_post_tags 
    ON posts USING gin (tags);  -- this is the correct way for arrays

-- Regular btree indexes (unchanged)
CREATE INDEX IF NOT EXISTS idx_users_username 
    ON users (username);

CREATE INDEX IF NOT EXISTS idx_posts_user_id 
    ON posts (user_id);

CREATE INDEX IF NOT EXISTS idx_comments_post_id 
    ON comments (post_id);