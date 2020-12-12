-- migrate:up
CREATE TYPE difficulty AS ENUM ('easy', 'medium', 'hard');
ALTER TABLE content_translations ADD COLUMN content_difficulty difficulty NOT NULL;

-- migrate:down
ALTER TABLE content_translations DROP COLUMN IF EXISTS content_difficulty;
DROP TYPE IF EXISTS difficulty;
