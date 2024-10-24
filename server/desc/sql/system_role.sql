CREATE TABLE `system_role` (
                               `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                               `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               `create_by` bigint DEFAULT NULL,
                               `update_by` bigint DEFAULT NULL,
                               `name` varchar(50) NOT NULL,
                               `status` varchar(1) NOT NULL,
                               `code` varchar(255) NOT NULL,
                               `desc` longtext NOT NULL,
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;