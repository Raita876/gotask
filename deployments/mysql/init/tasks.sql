CREATE DATABASE todoapp_db;

GRANT ALL ON todoapp_db.* TO mysql;

USE todoapp_db;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id CHAR(36) PRIMARY KEY NOT NULL UNIQUE,
    name CHAR(255) NOT NULL,
    email CHAR(255) NOT NULL,
    password CHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO
    users (id, name, email, password, created_at, updated_at)
VALUES
    (
        'b5e93ba4-ce33-4e36-83cb-c71177464a25',
        'hoge',
        'hoge@example.com',
        '$2a$10$wvjfqKN1W9q0McYB/CKrOeNy8z9wIjpYBbvhFu8VvbGxBD72Q1OL6', -- password
        '2024-09-10 10:00:00',
        '2024-09-10 10:00:00'
    );

DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks (
    id CHAR(36) PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(64) NOT NULL,
    status SMALLINT UNSIGNED NOT NULL,
    user_id CHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO
    tasks (id, name, status, user_id, created_at, updated_at)
VALUES
    (
        'b81240b0-7122-4d06-bdb2-8bcf512d6c63',
        'task1',
        0,
        'b5e93ba4-ce33-4e36-83cb-c71177464a25',
        '2024-09-10 10:00:00',
        '2024-09-10 10:00:00'
    ),
    (
        'fad796a1-e0ed-4ee5-9f88-9b7258d35ae9',
        'task2',
        1,
        'b5e93ba4-ce33-4e36-83cb-c71177464a25',
        '2024-09-13 10:00:00',
        '2024-09-13 10:00:00'
    ),
    (
        '07aaadbc-8967-406f-aebd-58b289377aef',
        'task3',
        2,
        'b5e93ba4-ce33-4e36-83cb-c71177464a25',
        '2024-09-12 10:00:00',
        '2024-09-12 10:00:00'
    );
