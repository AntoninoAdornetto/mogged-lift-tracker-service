CREATE TABLE `user` (
  `id` BINARY(16) PRIMARY KEY NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
  `first_name` VARCHAR(20) NOT NULL,
  `last_name` VARCHAR(20) NOT NULL,
  `email_address` VARCHAR(150) UNIQUE NOT NULL,
  `password` TEXT NOT NULL,
  `password_changed_at` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00'
);

CREATE TABLE `profile` (
	`id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`country` VARCHAR(5) NOT NULL,
	`measurement_system` VARCHAR(20) NOT NULL,
	`body_weight` REAL NOT NULL DEFAULT 0,
	`body_fat` REAL NOT NULL DEFAULT 0,
	`timezone` VARCHAR(50) NOT NULL,
	`user_id` BINARY(16) NOT NULL
);

CREATE TABLE `muscle_group` (
	`id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`name` VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE `category` (
	`id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`name` VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE `exercise` (
	`id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`name` VARCHAR(50) UNIQUE NOT NULL,
	`muscle_group` VARCHAR(20) NOT NULL,
	`category` VARCHAR(20) NOT NULL,
	`isStock` BOOL NOT NULL,
	`most_weight_lifted` REAL NOT NULL DEFAULT 0,
	`most_reps_lifted` SMALLINT UNSIGNED NOT NULL DEFAULT 0,
	`rest_timer` TIME NOT NULL DEFAULT '00:00:00',
	`user_id` BINARY(16) NOT NULL
);

CREATE TABLE `workout` (
  `id` BINARY(16) PRIMARY KEY NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
	`duration` TIME NOT NULL DEFAULT '00:00:00',
	`lifts` JSON,
	`user_id` BINARY(16) NOT NULL
);

CREATE TABLE `lift` (
  `id` BINARY(16) PRIMARY KEY NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
	`exercise_name` VARCHAR(50) NOT NULL,
	`weight_lifted` REAL NOT NULL,
	`reps` SMALLINT NOT NULL,
	`user_id` BINARY(16) NOT NULL,
	`workout_id` BINARY(16) NOT NULL
);

CREATE TABLE `template` (
	`id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`name` VARCHAR(50) UNIQUE NOT NULL,
	`lifts` JSON,
	`date_last_used` DATETIME NOT NULL DEFAULT NOW(),
	`created_by` BINARY(16) NOT NULL
);

CREATE INDEX `user_index_0` ON `user` (`id`);
CREATE INDEX `exercise_muscle_group_index_0` ON `exercise` (`muscle_group`);
CREATE INDEX `exercise_name_index_0` ON `exercise` (`name`);
CREATE INDEX `category_index_0` ON `exercise` (`category`);
CREATE INDEX `workout_user_index_0` ON `workout` (`user_id`);
CREATE INDEX `lift_user_index_0` ON `lift` (`user_id`);


ALTER TABLE `profile` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `exercise` ADD FOREIGN KEY (`muscle_group`) REFERENCES `muscle_group` (`name`);
ALTER TABLE `exercise` ADD FOREIGN KEY (`category`) REFERENCES `category` (`name`);
ALTER TABLE `exercise` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `workout` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `lift` ADD FOREIGN KEY (`exercise_name`) REFERENCES `exercise` (`name`);
ALTER TABLE `lift` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `lift` ADD FOREIGN KEY (`workout_id`) REFERENCES `workout` (`id`);

ALTER TABLE `template` ADD FOREIGN KEY (`created_by`) REFERENCES `user` (`id`);
