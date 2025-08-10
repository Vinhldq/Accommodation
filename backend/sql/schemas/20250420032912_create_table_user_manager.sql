-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_user_manager` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `account` VARCHAR(255) NOT NULL COMMENT 'account: email or SMS',
        `user_name` VARCHAR(255) NOT NULL DEFAULT "" COMMENT 'user name',
        `password` VARCHAR(60) NOT NULL COMMENT 'password',
        `login_time` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'login time',
        `logout_time` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'logout time',
        `login_ip` VARCHAR(45) NOT NULL DEFAULT "" COMMENT 'login IP',
        `is_deleted` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'is deleted: 0 - not deleted; 1 - deleted',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        UNIQUE KEY `unique_user_manager_account` (`account`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'user manager table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_user_manager`;

-- +goose StatementEnd