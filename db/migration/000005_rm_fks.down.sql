ALTER TABLE `profile` ADD FOREIGN KEY `profile_user_id_fk` (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

ALTER TABLE `exercise` ADD FOREIGN KEY `exercise_mg_name_fk` (`muscle_group`) REFERENCES `muscle_group` (`name`) ON DELETE CASCADE;
ALTER TABLE `exercise` ADD FOREIGN KEY `exercise_cat_name_fk` (`category`) REFERENCES `category` (`name`) ON DELETE CASCADE;
ALTER TABLE `exercise` ADD FOREIGN KEY `exercise_user_id_fk` (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

ALTER TABLE `workout` ADD FOREIGN KEY `wo_user_id_fk` (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;

ALTER TABLE `lift` ADD FOREIGN KEY `lift_exercise_name_fk` (`exercise_name`) REFERENCES `exercise` (`name`) ON DELETE CASCADE;
ALTER TABLE `lift` ADD FOREIGN KEY `lift_user_id_fk` (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;
ALTER TABLE `lift` ADD FOREIGN KEY `lift_wo_id_fk` (`workout_id`) REFERENCES `workout` (`id`) ON DELETE CASCADE;

ALTER TABLE `template` ADD FOREIGN KEY `template_user_id_fk` (`created_by`) REFERENCES `user` (`id`) ON DELETE CASCADE;

ALTER TABLE `session` ADD FOREIGN KEY `session_user_id_fk` (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE;
