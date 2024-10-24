CREATE TABLE `system_user` (
                               `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                               `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               `create_by` bigint DEFAULT NULL,
                               `update_by` bigint DEFAULT NULL,
                               `username` varchar(30) NOT NULL,
                               `password` varchar(100) NOT NULL,
                               `nickname` varchar(30) NOT NULL,
                               `gender` varchar(1) NOT NULL,
                               `phone` varchar(20) NOT NULL,
                               `email` varchar(100) NOT NULL,
                               `status` varchar(1) NOT NULL,
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;