CREATE TYPE GENDER AS ENUM ('male', 'female');

CREATE TYPE LANGUAGE_PREFERENCE AS ENUM ('en', 'ru');

CREATE TABLE IF NOT EXISTS users
(
    id                    UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    username              VARCHAR(255) NOT NULL UNIQUE,
    password_hash         VARCHAR(255) NOT NULL,
    salt                  VARCHAR(255) NOT NULL,
    email                 VARCHAR(255) NOT NULL UNIQUE,
    first_name            VARCHAR(255),
    last_name             VARCHAR(255),
    phone_number          VARCHAR(20),
    date_of_birth         DATE,
    gender                GENDER,
    bio                   TEXT,
    language_preference   LANGUAGE_PREFERENCE,
    created_at            TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    failed_login_attempts INT              DEFAULT 0,
    is_verified           BOOLEAN          DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS posts
(
    id         UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    user_id    UUID NOT NULL,
    content    TEXT NOT NULL,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    views      INT              DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_posts_id ON posts(id);

CREATE TABLE IF NOT EXISTS comments
(
    id         UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    post_id    UUID NOT NULL,
    user_id    UUID NOT NULL,
    content    TEXT NOT NULL,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE likes
(
    id         UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    post_id    UUID NOT NULL,
    user_id    UUID NOT NULL,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    UNIQUE (post_id, user_id)
);

CREATE TABLE Dislikes
(
    id         UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    post_id    UUID NOT NULL,
    user_id    UUID NOT NULL,
    created_at TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    UNIQUE (post_id, user_id)
);