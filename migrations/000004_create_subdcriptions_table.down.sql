DROP TABLE user_channels;
DROP TABLE user_following;

ALTER TABLE users ADD COLUMN channels JSONB NOT NULL DEFAULT '[]';
ALTER TABLE users ADD COLUMN following JSONB NOT NULL DEFAULT '[]';