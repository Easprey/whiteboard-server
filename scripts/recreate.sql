DROP TABLE fingerpoints;
DROP TABLE fingerpaths;
DROP TABLE users;

CREATE TABLE IF NOT EXISTS `whiteboard_dev`.`fingerpaths` (
  `path_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `board` VARCHAR(60) NOT NULL,
  `path_color` INT(11) UNSIGNED NULL DEFAULT NULL,
  `board_color` INT(11) UNSIGNED NULL DEFAULT NULL,
  `dash` TINYINT(4) NULL DEFAULT NULL,
  `blur` TINYINT(4) NULL DEFAULT NULL,
  `clear` TINYINT(4) NULL DEFAULT '0',
  `user_id` VARCHAR(45) NULL DEFAULT NULL,
  `stroke_width` INT(11) NULL DEFAULT NULL,
  PRIMARY KEY (`path_id`),
  INDEX `user_id_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `whiteboard_dev`.`users` (`user_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)


CREATE TABLE IF NOT EXISTS `whiteboard_dev`.`fingerpoints` (
  `path_id` INT(10) UNSIGNED NOT NULL,
  `point_num` SMALLINT(5) UNSIGNED NOT NULL AUTO_INCREMENT,
  `x` INT(10) UNSIGNED NULL DEFAULT NULL,
  `y` INT(10) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`point_num`),
  INDEX `path_id_idx` (`path_id` ASC) VISIBLE,
  CONSTRAINT `path_id`
    FOREIGN KEY (`path_id`)
    REFERENCES `whiteboard_dev`.`fingerpaths` (`path_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)

CREATE TABLE IF NOT EXISTS `whiteboard_dev`.`users` (
  `user_id` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`user_id`))
