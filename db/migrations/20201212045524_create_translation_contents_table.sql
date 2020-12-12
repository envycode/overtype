-- migrate:up

CREATE TYPE lang AS ENUM ('en', 'jp', 'jp-hiragana', 'jp-katakana');

CREATE TABLE IF NOT EXISTS content_translations (
    content_id uuid PRIMARY KEY NOT NULL,
    source_lang lang NOT NULL,
    destined_lang lang NOT NULL,
    source_text text NOT NULL,
    destined_text text NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now()
);

-- migrate:down

DROP TABLE IF EXISTS content_translations;
DROP TYPE IF EXISTS lang;

