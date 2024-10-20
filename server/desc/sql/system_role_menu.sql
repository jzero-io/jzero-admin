CREATE TABLE `system_role_menu` (
    `id` integer AUTO_INCREMENT,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `create_by` integer DEFAULT NULL,
    `update_by` integer DEFAULT NULL,
    `role_id` integer NOT NULL,
    `menu_id` integer NOT NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB COLLATE utf8mb4_general_ci;