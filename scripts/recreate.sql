DROP TABLE IF EXISTS fingerpoints;
DROP TABLE IF EXISTS fingerpaths;
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS `whiteboard_dev`.`users` (
  `user_id` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`user_id`));

CREATE TABLE IF NOT EXISTS `whiteboard_dev`.`fingerpaths` (
  `path_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `board` VARCHAR(60) NOT NULL,
  `path_color` INT UNSIGNED NULL DEFAULT NULL,
  `board_color` INT UNSIGNED DEFAULT NULL,
  `dash` BOOLEAN NULL DEFAULT NULL,
  `blur` BOOLEAN NULL DEFAULT NULL,
  `clear` BOOLEAN NULL DEFAULT '0',
  `user_id` VARCHAR(45) NULL DEFAULT NULL,
  `stroke_width` INT NULL DEFAULT NULL,
  PRIMARY KEY (`path_id`),
  INDEX `user_id_idx` (`user_id` ASC),
  CONSTRAINT `user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `whiteboard_dev`.`users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);

CREATE TABLE IF NOT EXISTS `whiteboard_dev`.`fingerpoints` (
  `point_num` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `path_id` INT UNSIGNED NOT NULL,
  `x` INT UNSIGNED NULL DEFAULT NULL,
  `y` INT UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`point_num`),
  INDEX `path_id_idx` (`path_id` ASC),
  CONSTRAINT `path_id`
    FOREIGN KEY (`path_id`)
    REFERENCES `whiteboard_dev`.`fingerpaths` (`path_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION);
