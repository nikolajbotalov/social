CREATE TABLE auth (
    id UUID PRIMARY KEY,
    nickname VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    FOREIGN KEY (id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_auth_nickname ON auth(nickname);