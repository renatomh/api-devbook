CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    username varchar(50) not null unique,
    email varchar(50) not null unique,
    pass varchar(100) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;
