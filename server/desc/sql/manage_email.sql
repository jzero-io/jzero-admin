CREATE TABLE `manage_email` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_by` bigint DEFAULT NULL,
  `update_by` bigint DEFAULT NULL,
  `from` longtext NOT NULL,
  `host` longtext NOT NULL,
  `port` bigint NOT NULL,
  `username` longtext NOT NULL,
  `password` longtext NOT NULL,
  `enable_ssl` tinyint(1) NOT NULL,
  `is_verify` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci