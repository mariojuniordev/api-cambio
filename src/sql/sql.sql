CREATE DATABASE IF NOT EXISTS apiCambio;
USE apiCambio;

DROP TABLE IF EXISTS depositos;

CREATE TABLE depositos(
  id INT auto_increment PRIMARY KEY,
  valorDeposito FLOAT NOT NULL  
) ENGINE=INNODB;