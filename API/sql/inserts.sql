insert into users (name, nick, email, password)
values
("user1", "youser_1", "user_1@gmail.com", "$2a$10$895nvMPssrff/CzIYoNoOeNmL4EPNC9TkYAxh/J4y5NSsS291E3XK"),
("user2", "youser_2", "user_2@gmail.com", "$2a$10$895nvMPssrff/CzIYoNoOeNmL4EPNC9TkYAxh/J4y5NSsS291E3XK"),
("user3", "youser_3", "user_3@gmail.com", "$2a$10$895nvMPssrff/CzIYoNoOeNmL4EPNC9TkYAxh/J4y5NSsS291E3XK");

insert into followers (ID_user, ID_follower)
values
(1, 2),
(1, 3),
(3, 1),
(3, 2),
(2, 1);