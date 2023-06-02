CREATE TRIGGER `inactive` AFTER DELETE ON user
	FOR EACH ROW
		INSERT INTO `inactive_user` (id) VALUES (OLD.id);
