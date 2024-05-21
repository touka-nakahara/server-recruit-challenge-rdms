CREATE DATABASE IF NOT EXISTS music;

CREATE TABLE IF NOT EXISTS music.singer(
  id int unsigned not null auto_increment, 
  name  varchar(255) not null,
  primary key (id)
);
