CREATE TABLE `system_user` (
    `id` integer AUTO_INCREMENT,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `create_by` integer DEFAULT NULL,
    `update_by` integer DEFAULT NULL,
    `username` varchar(30) NOT NULL DEFAULT '',
    `password` varchar(100) NOT NULL DEFAULT '',
    `nickname` varchar(30) NOT NULL DEFAULT '',
    `gender` varchar(1) NOT NULL DEFAULT '0',
    `phone` varchar(20) NOT NULL DEFAULT '',
    `email` varchar(100) NOT NULL DEFAULT '',
    `status` varchar(1) NOT NULL DEFAULT '0',
    PRIMARY KEY (id)
)ENGINE = InnoDB COLLATE utf8mb4_general_ci;