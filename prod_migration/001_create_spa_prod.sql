-- +goose Up
CREATE TABLE IF NOT EXISTS posts
(
    id          uuid      not null,
    title       text      not null,
    created     timestamp not null,
    description text      not null default '',
    user_id     uuid      not null
);

CREATE TABLE IF NOT EXISTS users
(
    id        uuid    not null,
    username  text    not null,
    password  text    not null
);

CREATE TABLE IF NOT EXISTS comments
(
    id          uuid    not null,
    username    text    not null,
    content     text    not null,
    user_id     uuid    not null,
    post_id     uuid    not null
);

-- +goose Down
-- DROP TABLE spa-test;