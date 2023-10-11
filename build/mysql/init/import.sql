CREATE DATABASE IF NOT EXISTS `crud_database`;
USE `crud_database`;

CREATE TABLE IF NOT EXISTS `user` (
    `username` varchar(255) NOT NULL,
    `password` int NOT NULL,
    PRIMARY KEY (`username`)
    );