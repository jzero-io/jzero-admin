CREATE TABLE `system_role` (
    `id` integer AUTO_INCREMENT,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `create_by` integer DEFAULT NULL,
    `update_by` integer DEFAULT NULL,
    `name` varchar(50) NOT NULL,
    `status` varchar(1) NOT NULL,
    `code` varchar(255) NOT NULL,
    `desc` text NOT NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB COLLATE utf8mb4_general_ci;