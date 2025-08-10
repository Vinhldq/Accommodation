-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS `ecommerce_go_order` (
        `id` VARCHAR(36) NOT NULL COMMENT 'ID',
        `user_id` VARCHAR(36) NOT NULL COMMENT 'user base ID',
        `final_total` DECIMAL(15, 0) NOT NULL COMMENT 'final total',
        -- `original_total` DECIMAL(10, 2) NOT NULL COMMENT 'original total',
        `order_id_external` VARCHAR(8) NOT NULL COMMENT 'order id external',
        `order_status` ENUM (
            'pending_payment',
            'payment_failed',
            'payment_success',
            'checked_in',
            'completed',
            'canceled',
            'refunded'
        ) NOT NULL COMMENT "order status",
        `accommodation_id` VARCHAR(36) NOT NULL COMMENT 'accommodation ID',
        `voucher_id` VARCHAR(36) COMMENT 'voucher ID',
        `checkin_date` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `checkout_date` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        `created_at` BIGINT UNSIGNED NOT NULL COMMENT 'created at',
        `updated_at` BIGINT UNSIGNED NOT NULL COMMENT 'updated at',
        PRIMARY KEY (`id`),
        INDEX `idx_order_status_checkin_checkout` (`order_status`, `checkin_date`, `checkout_date`),
        INDEX `idx_order_accommodation_id` (`accommodation_id`),
        FOREIGN KEY (`user_id`) REFERENCES `ecommerce_go_user_base` (`id`) ON DELETE CASCADE,
        FOREIGN KEY (`voucher_id`) REFERENCES `ecommerce_go_voucher` (`id`) ON DELETE CASCADE,
        FOREIGN KEY (`accommodation_id`) REFERENCES `ecommerce_go_accommodation` (`id`) ON DELETE CASCADE
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'order table';

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ecommerce_go_order`;

-- +goose StatementEnd