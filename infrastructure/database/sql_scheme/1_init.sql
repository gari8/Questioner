DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS questions;
DROP TABLE IF EXISTS answers;
DROP TABLE IF EXISTS choices;
DROP TABLE IF EXISTS choice_answers;

CREATE TABLE IF NOT EXISTS users
(
    id TEXT NOT NULL,
    username TEXT NOT NULL,
    icon TEXT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    description TEXT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS questions
(
    id TEXT NOT NULL,
    title TEXT NOT NULL,
    user_id TEXT NOT NULL,
    content TEXT NOT NULL,
    published BOOLEAN NOT NULL DEFAULT TRUE,
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    answer_type INT NOT NULL,
    published_at TIMESTAMP NULL,
    finished_at TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS answers
(
    id SERIAL NOT NULL,
    user_id TEXT NOT NULL,
    question_id TEXT NOT NULL,
    content TEXT NOT NULL,
    answer_type INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS choices
(
    id SERIAL NOT NULL,
    question_id TEXT NOT NULL,
    content TEXT NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS choice_answers
(
    id SERIAL NOT NULL,
    user_id TEXT NOT NULL,
    choice_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
);




INSERT INTO users(id, username, email, password, created_at) VALUES ('12345678', 'taro', 'taro@gmail.com', 'tarosugoiyo', );

