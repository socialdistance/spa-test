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

-- CREATE UNIQUE INDEX posts_id ON posts(id);

INSERT INTO users (id, username, password) VALUES
('49842fff-5c94-4eca-8920-d41e3e7bc61c', 'test1', 'test1'),
('74de454f-ee87-4e13-8bc7-85e7662ca842', 'test2', 'test2'),
('35161795-b7c6-404c-ae11-8056f1585fa9', 'test3', 'test3'),
('d3bdc536-7386-4ba2-9143-f71503ab551c', 'test4', 'test4'),
('00eb65bc-214d-4050-b404-5857d2412041', 'test5', 'test5'),
('bb16a31f-4e8c-43ee-bf79-39021cce11ea', 'test6', 'test6'),
('7ac9b0af-9901-4bcc-863c-d8f8e79abdf2', 'test7', 'test7'),
('6665e2c0-d93e-4c7c-aad3-d05aae176b5f', 'test8', 'test8'),
('cf57e23b-5fa7-4d5d-8287-7af062a9b7cc', 'test9', 'test9'),
('cab71307-4d18-4c52-804c-d5be36a873c5', 'test10', 'test10')
;

INSERT INTO comments (id, username, content, user_id, post_id) VALUES
('3cdec67f-2265-4cfa-af5a-3b7dd3bf04bd', 'test1 comment1', 'comment1', '49842fff-5c94-4eca-8920-d41e3e7bc61c', 'fcfa069c-28a9-48d6-b48d-befc8133f2b4'),
('8c32b6c3-06be-4e7a-9007-b80d041d916c', 'test1 comment2', 'comment2', '74de454f-ee87-4e13-8bc7-85e7662ca842', 'fcfa069c-28a9-48d6-b48d-befc8133f2b4'),
('881a6bf4-7d10-4cc5-a3a5-00e42e9c4bfd', 'test2 comment1', 'comment3', '35161795-b7c6-404c-ae11-8056f1585fa9', '5753e882-91e0-4e1a-a827-eef8d8271e50')
;

INSERT INTO posts (id, title, created, description, user_id) VALUES
('fcfa069c-28a9-48d6-b48d-befc8133f2b4', 'Post 1', '2022-01-06 01:09:00', 'Post description 1', 'f3bed4b6-8695-47f2-a5c3-3a0509e7f8a6'),
('87a846ce-4488-45a9-bd36-ca72368a7185', 'Post 2', '2022-01-06 01:10:00', 'Post description 2', '4ca38f2a-03bc-412a-9a8e-c3d96dd6fcb6'),
('5753e882-91e0-4e1a-a827-eef8d8271e50', 'Post 3', '2022-01-06 02:45:00', 'Post description 3', '2673819d-74d2-4bae-bc0a-f6e1956bc4b3'),
('fcfa069c-28a9-48d6-b48d-befc8133f2b2', 'Post 4', '2022-01-06 03:00:00', 'Post description 4', 'f1bed4b6-8695-47f2-a5c3-3a0509e7f8a6'),
('87a876ce-4488-45a9-bd36-ca72368a7181', 'Post 5', '2022-01-06 04:00:00', 'Post description 5', '2ca38f2a-03bc-412a-9a8e-c3d96dd6fcb6'),
('5753e882-91e0-4e1a-a827-eef8d8271e57', 'Post 6', '2022-01-06 05:45:00', 'Post description 6', '3673819d-74d2-4bae-bc0a-f6e1956bc4b3'),
('fcfa069c-28a9-48d6-b48d-befc8133f3b4', 'Post 7', '2022-01-06 06:00:00', 'Post description 7', 'f5bed4b6-8695-47f2-a5c3-3a0509e7f8a6'),
('87a876ce-4488-45a9-bd36-ca72368a7285', 'Post 8', '2022-01-06 07:00:00', 'Post description 8', '4ca78f2a-03bc-412a-9a8e-c3d96dd6fcb6'),
('4753e882-91e0-4e1a-a827-eef8d8271e50', 'Post 9', '2022-01-06 08:45:00', 'Post description 9', '2672819d-74d2-4bae-bc0a-f6e1956bc4b3'),
('fcfa069c-28a9-48d6-b48d-befc8133f2b2', 'Post 10', '2022-01-06 09:00:00', 'Post description 10', 'f8bed4b6-8695-47f2-a5c3-3a0509e7f8a6'),
('87a876ce-4488-45a9-bd36-ca72368a7185', 'Post 11', '2022-01-06 10:00:00', 'Post description 11', '4ca98f2a-03bc-412a-9a8e-c3d96dd6fcb6'),
('5753e882-91e0-4e1a-a827-eef8d1271e50', 'Post 12', '2022-01-06 11:45:00', 'Post description 12', '2623819d-74d2-4bae-bc0a-f6e1956bc4b3'),
('fcfa069c-28a9-48d6-b48d-befc8123f2b4', 'Post 13', '2022-01-06 12:00:00', 'Post description 13', 'f3bed6b6-8695-47f2-a5c3-3a0509e7f8a6'),
('87a876ce-4488-45a9-bd36-ca72348a7185', 'Post 14', '2022-01-06 13:00:00', 'Post description 14', '4ca38f4a-03bc-412a-9a8e-c3d96dd6fcb6'),
('5753e882-91e0-4e1a-a827-eef8e8271e50', 'Post 15', '2022-01-06 14:45:00', 'Post description 15', '2673818d-74d2-4bae-bc0a-f6e1956bc4b3'),
('fcfa069c-28a9-42d6-b48d-befc8133f2b4', 'Post 16', '2022-01-06 15:00:00', 'Post description 16', 'f3bed4b6-8495-47f2-a5c3-3a0509e7f8a6'),
('87a876ce-4488-43a9-bd36-ca72368a7185', 'Post 17', '2022-01-06 16:00:00', 'Post description 17', '4ca38f2a-02bc-412a-9a8e-c3d96dd6fcb6'),
('5753e882-01e0-4e5a-a827-eef8d8271e50', 'Post 18', '2022-01-06 17:45:00', 'Post description 18', '2673819d-74d2-4bae-bc0a-f6e1956bc4b3'),
('fcfa069c-28a9-18d6-b48d-befc8133f2b2', 'Post 19', '2022-01-06 18:00:00', 'Post description 19', 'f1bed4b6-8195-47f2-a5c3-3a0509e7f8a6'),
('87a876ce-4488-35a9-bd36-ca72368a7181', 'Post 20', '2022-01-06 19:00:00', 'Post description 20', '2ca38f2a-02bc-452a-9a8e-c3d96dd6fcb6'),
('5753e882-91e0-8e1a-a827-eef8d8271e57', 'Post 21', '2022-01-06 20:45:00', 'Post description 21', '3673819d-73d2-4bae-bc0a-f6e1956bc4b3'),
('fcfa069c-28a9-28d6-b48d-befc8133f3b4', 'Post 22', '2022-01-06 21:00:00', 'Post description 22', 'f5bed4b6-8495-47f2-a5c3-3a0509e7f8a6'),
('87a876ce-4488-15a9-bd36-ca72368a7285', 'Post 23', '2022-01-06 22:00:00', 'Post description 23', '4ca78f2a-02bc-452a-9a8e-c3d96dd6fcb6'),
('5753e882-91e0-6e1a-a827-eef8d8271e50', 'Post 24', '2022-01-06 23:45:00', 'Post description 24', '2672819d-77d2-4bae-bc0a-f6e1956bc4b3'),
('fcfa069c-28a9-78d6-b48d-befc8133f2b4', 'Post 25', '2022-01-06 24:00:00', 'Post description 25', 'f8bed4b6-8395-42f2-a5c3-3a0509e7f8a6'),
('87a876ce-4488-35a9-bd36-ca72368a7185', 'Post 26', '2022-02-06 09:00:00', 'Post description 26', '4ca98f2a-02bc-212a-9a8e-c3d96dd6fcb6'),
('5753e882-91e0-1e1a-a827-eef8d1271e50', 'Post 27', '2022-02-06 10:45:00', 'Post description 27', '2623819d-71d2-5bae-bc0a-f6e1956bc4b3'),
('fcfa069c-28a9-38d6-b48d-befc8123f2b4', 'Post 28', '2022-02-06 11:00:00', 'Post description 28', 'f3bed6b6-8695-27f2-a5c3-3a0509e7f8a6'),
('87a876ce-4488-85a9-bd36-ca72348a7185', 'Post 29', '2022-02-06 12:00:00', 'Post description 29', '4ca38f4a-03bc-492a-9a8e-c3d96dd6fcb6'),
('5753e882-91e0-6e1a-a827-eef8e8271e50', 'Post 30', '2022-02-06 13:45:00', 'Post description 30', '2673818d-74d2-2bae-bc0a-f6e1956bc4b3')
;
-- +goose Down
-- DROP TABLE spa-test;