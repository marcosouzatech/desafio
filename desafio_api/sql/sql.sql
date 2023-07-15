CREATE DATABASE IF NOT EXISTS brasil;
USE brasil;

DROP TABLE IF EXISTS itens; 

CREATE TABLE IF NOT EXISTS itens(
    id int auto_increment primary key, 
    Product varchar(10) not null, 
    Nome varchar(50) not null unique,
    Categoria varchar(50) not null unique,
    Token varchar(20) not null unique,
    CriadoEm timeStamp default current_timestamp()
) ENGINE=INNODB;