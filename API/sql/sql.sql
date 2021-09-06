CREATE DATABASE IF NOT EXISTS golang;
USE golang;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    createdOn timestamp default current_timestamp()
 ) ENGINE=INNODB;