CREATE DATABASE IF NOT EXISTS golang;
USE golang;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;
    
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
