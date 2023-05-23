CREATE TABLE `session` (
	`id` BINARY(16) PRIMARY KEY DEFAULT (UUID_TO_BIN(UUID())),
	`refresh_token` TEXT NOT NULL,
	`user_agent` TEXT NOT NULL,
	`client_ip` TEXT NOT NULL,
	`is_banned` BOOL NOT NULL DEFAULT FALSE,
	`expires_at` TIMESTAMP NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
	`user_id` BINARY(16) NOT NULL
);

ALTER TABLE `session` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;