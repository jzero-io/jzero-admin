CREATE TABLE `system_user_role` (
    `id` integer AUTO_INCREMENT,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `create_by` integer DEFAULT NULL,
    `update_by` integer DEFAULT NULL,
    `user_id` integer NOT NULL DEFAULT 0,
    `role_id` integer NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
)ENGINE = InnoDB COLLATE utf8mb4_general_ci;