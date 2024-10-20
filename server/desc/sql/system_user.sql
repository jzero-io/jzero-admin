CREATE TABLE `system_user` (
    `id` integer AUTO_INCREMENT,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `create_by` integer DEFAULT NULL,
    `update_by` integer DEFAULT NULL,
    `username` varchar(30) NOT NULL,
    `password` varchar(100) NOT NULL,
    `nickname` varchar(30) NOT NULL,
    `gender` varchar(1) NOT NULL,
    `phone` varchar(20) NOT NULL,
    `email` varchar(100) NOT NULL,
    `status` varchar(1) NOT NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB COLLATE utf8mb4_general_ci;