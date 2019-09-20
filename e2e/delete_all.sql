DELETE FROM fingerpoints;
DELETE FROM fingerpaths;
DELETE FROM users;

ALTER TABLE fingerpoints AUTO_INCREMENT = 1;
ALTER TABLE fingerpaths AUTO_INCREMENT = 1;