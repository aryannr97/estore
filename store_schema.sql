CREATE TABLE IF NOT EXISTS `products` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(255),
	`seller` VARCHAR(255),
	`quantity` INT,
	`price` INT,
	PRIMARY KEY (`id`)
);

INSERT INTO `products`(`name`,`seller`,`quantity`,`price`) VALUES("shoes","nike",10,2000);

CREATE TABLE IF NOT EXISTS `cart` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`userid` INT,
	`productid` INT,
	`quantity` INT,
	`price` INT,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `reserved_products` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`checkoutid` VARCHAR(255),
	`productid` INT,
	`quantity` INT,
	PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `orders` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`userid` INT,
	`txid` VARCHAR(255),
	`checkoutid` VARCHAR(255),
	`detail` VARCHAR(512),
	PRIMARY KEY (`id`)
);