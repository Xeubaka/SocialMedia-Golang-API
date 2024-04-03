insert into users (name, nick, email, password)
values
("User 1", "user_1", "user_1@gmail.com", "$2a$10$jINWavIVNsSjGNYZSbKhnuaPpMt68e7Bxa6iMRI3ILXpYU2eU5r56"),
("User 2", "user_2", "user_2@gmail.com", "$2a$10$jINWavIVNsSjGNYZSbKhnuaPpMt68e7Bxa6iMRI3ILXpYU2eU5r56"),
("User 3", "user_3", "user_3@gmail.com", "$2a$10$jINWavIVNsSjGNYZSbKhnuaPpMt68e7Bxa6iMRI3ILXpYU2eU5r56");

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into posts(title, content, author_id)
values
("Post of user 1", "this is the post of user 1! Yay!", 1),
("Post of user 2", "this is the post of user 2! Yay!", 2),
("Post of user 2", "this is the post of user 3! Yay!", 3);