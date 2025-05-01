-- +goose Up
-- +goose StatementBegin
CREATE TABLE `feedbacks` (
    `id` varchar(36) NOT NULL,
    `nursing_id` varchar(36) NOT NULL,
    `medical_record_id` varchar(36) NOT NULL,
    `patient_name` varchar(70),
    `service` text,
    `star` enum('1', '2', '3', '4', '5') DEFAULT '5',
    `content` text,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `feedbacks`;
-- +goose StatementEnd
