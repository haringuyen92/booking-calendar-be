-- Modify "courses" table
ALTER TABLE `courses` MODIFY COLUMN `active` bigint unsigned NULL, ADD COLUMN `color` longtext NULL;
