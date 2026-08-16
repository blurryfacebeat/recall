-- +goose Up

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuidv7(),

                       telegram_id BIGINT UNIQUE,
                       telegram_username TEXT,

                       display_name TEXT NOT NULL,
                       avatar_url TEXT,

                       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down

DROP TABLE users;