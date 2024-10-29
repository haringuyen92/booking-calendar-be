-- Create "stores" table
CREATE TABLE `stores` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `user_id` bigint unsigned NULL,
  `name` varchar(255) NULL,
  `description` text NULL,
  `address` varchar(255) NULL,
  `website` varchar(255) NULL,
  `logo` varchar(255) NULL,
  `email` varchar(255) NULL,
  `phone` varchar(255) NULL,
  `location` varchar(255) NULL,
  `status` bigint unsigned NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
