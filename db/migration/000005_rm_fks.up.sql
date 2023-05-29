-- planet scale does not support fks.
ALTER TABLE `profile` DROP FOREIGN KEY `profile_ibfk_1`;

ALTER TABLE `exercise` DROP FOREIGN KEY `exercise_ibfk_1`;
ALTER TABLE `exercise` DROP FOREIGN KEY `exercise_ibfk_2`;
ALTER TABLE `exercise` DROP FOREIGN KEY `exercise_ibfk_3`;

ALTER TABLE `workout` DROP FOREIGN KEY `workout_ibfk_1`;

ALTER TABLE `lift` DROP FOREIGN KEY `lift_ibfk_1`;
ALTER TABLE `lift` DROP FOREIGN KEY `lift_ibfk_2`;
ALTER TABLE `lift` DROP FOREIGN KEY `lift_ibfk_3`;

ALTER TABLE `template` DROP FOREIGN KEY `template_ibfk_1`;

ALTER TABLE `session` ADD FOREIGN KEY `session_ibfk_1` (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

