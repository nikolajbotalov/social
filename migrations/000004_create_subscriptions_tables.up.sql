-- Подписки на каналы
CREATE TABLE user_channels (
    user_id UUID NOT NULL,
    channel_id UUID NOT NULL,
    PRIMARY KEY (user_id, channel_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE
);

-- Подписки на пользователей
CREATE TABLE user_following (
    follower_id UUID NOT NULL,
    following_id UUID NOT NULL,
    PRIMARY KEY (follower_id, following_id),
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (following_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Удаляем старые JSONB-поля из users
ALTER TABLE users DROP COLUMN channels;
ALTER TABLE users DROP COLUMN following;