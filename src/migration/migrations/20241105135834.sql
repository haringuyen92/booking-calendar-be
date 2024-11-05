-- Modify "courses" table
ALTER TABLE `courses` ADD COLUMN `store_id` bigint unsigned NULL, ADD COLUMN `image` longtext NULL, ADD COLUMN `description` longtext NULL, ADD COLUMN `cost` bigint unsigned NULL, ADD COLUMN `estimate_time` bigint unsigned NULL, ADD COLUMN `active` bool NULL, ADD COLUMN `position` bigint unsigned NULL;
-- Modify "staffs" table
ALTER TABLE `staffs` RENAME COLUMN `user_id` TO `store_id`, MODIFY COLUMN `position` bigint unsigned NULL, ADD COLUMN `phone` longtext NULL, ADD COLUMN `cost` bigint unsigned NULL, ADD COLUMN `max_booking_slot` bigint unsigned NULL, ADD COLUMN `active` bool NULL, ADD COLUMN `color` longtext NULL, ADD COLUMN `is_all_course` bool NULL;
