CREATE TABLE `user` (
  `id` BINARY(16) PRIMARY KEY NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
  `first_name` VARCHAR(20) NOT NULL,
  `last_name` VARCHAR(20) NOT NULL,
  `email_address` TEXT NOT NULL,
  `password` TEXT NOT NULL,
  `password_changed_at` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00'
);

CREATE TABLE `profile` (
	`id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`country` VARCHAR(5) NOT NULL,
	`measurement_system` VARCHAR(20) NOT NULL,
	`body_weight` REAL NOT NULL DEFAULT 0,
	`body_fat` REAL NOT NULL DEFAULT 0,
	`user_id` BINARY(16) NOT NULL
);
