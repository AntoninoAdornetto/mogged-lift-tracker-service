CREATE TABLE `inactive_users` (
	`id` BINARY(16) PRIMARY KEY,
	`recorded` DATE DEFAULT (NOW()),
	`complete` BOOLEAN DEFAULT false
);

CREATE TRIGGER `inactive` AFTER DELETE ON user
	FOR EACH ROW
		INSERT INTO `inactive_users` (id) VALUES (OLD.id);
