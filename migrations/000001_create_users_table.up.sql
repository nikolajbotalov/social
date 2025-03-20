CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    nickname VARCHAR(255) NOT NULL UNIQUE,
    birthday DATE,
    last_visit TIMESTAMP WITH TIME ZONE NOT NULL,
    interests JSONB NOT NULL DEFAULT '[]',
    channels JSONB NOT NULL DEFAULT '[]',
    following JSONB NOT NULL DEFAULT '[]',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX idx_users_nickname ON users(nickname);