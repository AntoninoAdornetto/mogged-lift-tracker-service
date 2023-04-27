CREATE TABLE `user` (
	`id` BINARY(16) PRIMARY KEY NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
	`first_name` VARCHAR(20) NOT NULL,
	`last_name` VARCHAR(20) NOT NULL,
	`email_address` VARCHAR(150) UNIQUE NOT NULL,
	`password` TEXT NOT NULL,
	`password_changed_at` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00'
);

CREATE TABLE `profile` (
	`id` MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
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

-- using TIME for rest_timer results in a typing error with sqlc when using a default value of '00:00:00'
-- we can store as a string and later parse it as a duration using go's parseDuration method
CREATE TABLE `exercise` (
	`id` MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`name` VARCHAR(50) UNIQUE NOT NULL,
	`muscle_group` VARCHAR(20) NOT NULL,
	`category` VARCHAR(20) NOT NULL,
	`isStock` BOOL NOT NULL DEFAULT false,
	`most_weight_lifted` REAL NOT NULL DEFAULT 0,
	`most_reps_lifted` SMALLINT UNSIGNED NOT NULL DEFAULT 0,
	`rest_timer` VARCHAR(15) NOT NULL DEFAULT '00:00:00s', 
	`user_id` BINARY(16) NOT NULL
);

-- using TIME for duration results in a typing error with sqlc when using a default value of '00:00:00'
-- we can store as a string and later parse it as a duration using go's parseDuration method
CREATE TABLE `workout` (
	`id` MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`duration` VARCHAR(10) NOT NULL DEFAULT '00:00:00s',
	`lifts` JSON,
	`user_id` BINARY(16) NOT NULL 
);

CREATE TABLE `lift` (
  `id` SERIAL PRIMARY KEY,
	`exercise_name` VARCHAR(50) NOT NULL,
	`weight_lifted` REAL NOT NULL,
	`reps` SMALLINT NOT NULL,
	`set_type` VARCHAR(25) NOT NULL,
	`user_id` BINARY(16) NOT NULL,
	`workout_id` MEDIUMINT UNSIGNED NOT NULL 
);

CREATE TABLE `template` (
	`id` MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`name` VARCHAR(50) UNIQUE NOT NULL,
	`lifts` JSON,
	`date_last_used` DATE NOT NULL DEFAULT '1900-01-01',
	`created_by` BINARY(16) NOT NULL
);

CREATE INDEX `user_index_0` ON `user` (`id`);
CREATE INDEX `exercise_muscle_group_index_0` ON `exercise` (`muscle_group`);
CREATE INDEX `user_exercises_index_0` ON `exercise` (`user_id`);
CREATE INDEX `exercise_name_index_0` ON `exercise` (`name`);
CREATE INDEX `category_index_0` ON `exercise` (`category`);
CREATE INDEX `workout_user_index_0` ON `workout` (`user_id`);
CREATE INDEX `lift_user_index_0` ON `lift` (`user_id`);


ALTER TABLE `profile` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

ALTER TABLE `exercise` ADD FOREIGN KEY (`muscle_group`) REFERENCES `muscle_group` (`name`);
ALTER TABLE `exercise` ADD FOREIGN KEY (`category`) REFERENCES `category` (`name`);
ALTER TABLE `exercise` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

ALTER TABLE `workout` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

ALTER TABLE `lift` ADD FOREIGN KEY (`exercise_name`) REFERENCES `exercise` (`name`);
ALTER TABLE `lift` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;
ALTER TABLE `lift` ADD FOREIGN KEY (`workout_id`) REFERENCES `workout` (`id`) ON DELETE CASCADE;
-- @TODO: determine if index on lift.workout_id can be beneficial

ALTER TABLE `template` ADD FOREIGN KEY (`created_by`) REFERENCES `user` (`id`) ON DELETE CASCADE;
