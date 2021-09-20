CREATE DATABASE IF NOT EXISTS golang;
USE golang;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

    
CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createdOn timestamp default current_timestamp()
 ) ENGINE=INNODB;

 CREATE TABLE followers(
    ID_user int not null,
    FOREIGN KEY (ID_user)
    REFERENCES users(id)
    ON DELETE CASCADE,

    ID_follower int not null,
    FOREIGN KEY (ID_follower)
    REFERENCES users(id)
    ON DELETE CASCADE,

    primary key (ID_user, ID_follower)
 ) ENGINE=INNODB;

 CREATE TABLE posts(
   id int auto_increment primary key,
   title varchar(50) not null,   
   content varchar(300) not null,
   
   poster_id int not null,
   FOREIGN KEY (poster_id)
   REFERENCES users(id)
   ON DELETE CASCADE,   

   likes int default 0,
   createdOn timestamp default current_timestamp()
 ) ENGINE=INNODB;
