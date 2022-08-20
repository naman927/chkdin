-- Active: 1660978741214@@127.0.0.1@3306@demo
create table if not EXISTS users(
  id int not null key auto_increment,
  name varchar(255),
  password varchar(255),
  username varchar(255) UNIQUE,
  token varchar(255),
  expire_at TIMESTAMP
);