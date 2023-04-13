CREATE TABLE `user` (
  `id` BINARY(36) PRIMARY KEY NOT NULL,
  `first_name` VARCHAR(20) NOT NULL,
  `last_name` VARCHAR(20) NOT NULL,
  `email_address` TEXT NOT NULL,
  `password` TEXT NOT NULL,
  `password_changed_at` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00'
);

CREATE TABLE `profile` (
  `id` BINARY(36) PRIMARY KEY NOT NULL,
  `country` VARCHAR(255) NOT NULL,
  `measurement_system` VARCHAR(20) NOT NULL,
  `body_weight` REAL,
  `body_fat` REAL,
  `user_id` BINARY(36) NOT NULL
);

CREATE TABLE `muscle_group` (
  `id` SMALLINT PRIMARY KEY,
  `name` VARCHAR(255) UNIQUE NOT NULL
);


CREATE TABLE `category` (
  `id` SMALLINT PRIMARY KEY,
  `name` VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE `exercise` (
  `id` SERIAL PRIMARY KEY,
  `name` VARCHAR(255) UNIQUE NOT NULL,
  `muscle_group` VARCHAR(255) NOT NULL,
  `category` VARCHAR(255) NOT NULL,
  `isStock` BOOL NOT NULL,
  `most_weight_lifted` REAL DEFAULT 0,
  `most_reps_lifted` SMALLINT DEFAULT 0,
  `rest_timer` TIME NOT NULL DEFAULT '00:00:00',
  `user_id` BINARY(36) 
);

CREATE TABLE `workout` (
  `id` BINARY(36) PRIMARY KEY,
  `start_time` TIME,
  `finish_time` TIME,
  `lifts` JSON,
  `user_id` BINARY(36) 
);

CREATE TABLE `lift` (
  `id` BINARY(36) PRIMARY KEY,
  `exercise_name` VARCHAR(255) NOT NULL,
  `weight_lifted` REAL NOT NULL,
  `reps` SMALLINT NOT NULL,
  `user_id` BINARY(36),
  `workout_id` BINARY(36)
);

CREATE TABLE `template` (
  `id` BINARY(36) PRIMARY KEY,
  `name` VARCHAR(255) UNIQUE,
  `lifts` JSON,
  `muscle_group_category` VARCHAR(255) NOT NULL,
  `date_last_used` DATETIME NOT NULL DEFAULT '1970-01-01 00:00:00',
  `created_by` BINARY(36) NOT NULL
);

CREATE INDEX `user_index_0` ON `user` (`id`);

CREATE INDEX `muscle_group_index_0` ON exercise (`muscle_group`);

ALTER TABLE `exercise` ADD INDEX `category_index` (`category`);

ALTER TABLE `exercise` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `workout` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `lift` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `lift` ADD FOREIGN KEY (`workout_id`) REFERENCES `workout` (`id`);

ALTER TABLE `profile` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

ALTER TABLE `muscle_group` ADD FOREIGN KEY (`name`) REFERENCES `exercise` (`muscle_group`);

ALTER TABLE `category` ADD FOREIGN KEY (`name`) REFERENCES `exercise` (`category`);

ALTER TABLE `template` ADD FOREIGN KEY (`muscle_group_category`) REFERENCES `muscle_group` (`name`);

ALTER TABLE `template` ADD FOREIGN KEY (`created_by`) REFERENCES `user` (`id`);
