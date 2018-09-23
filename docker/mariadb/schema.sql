CREATE TABLE `order` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `seller_id` INT NULL,
  `customer_id` INT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `seller` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(60) NULL,
  `last_name` VARCHAR(65) NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `customer` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(60) NULL,
  `last_name` VARCHAR(65) NULL,
  `email` VARCHAR(95) NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `phone` (
  `number` VARCHAR(30) NOT NULL,
  `user_id` INT NOT NULL);

CREATE TABLE `address` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `street_address` VARCHAR(255) NULL,
  `city` VARCHAR(45) NULL,
  `zip_code` VARCHAR(20) NULL,
  `user_id` INT NOT NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `pseudo` VARCHAR(45) NOT NULL,
  `email` VARCHAR(45) NOT NULL,
  `isAdmin` TINYINT NOT NULL DEFAULT 0,
  `password` VARCHAR(65) NOT NULL,
  PRIMARY KEY (`id`));
