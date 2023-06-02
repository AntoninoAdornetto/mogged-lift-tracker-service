CREATE TABLE `inactive_user` (
	`id` BINARY(16) PRIMARY KEY,
	`recorded` DATETIME DEFAULT (NOW()) 
);
