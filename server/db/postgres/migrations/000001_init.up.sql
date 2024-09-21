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