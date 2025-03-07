-- +goose Up
-- +goose StatementBegin
CREATE TABLE `nursing` (
    `id` varchar(36) NOT NULL,
    `nurse_picture` varchar(300),
    `nurse_name` varchar(70),
    `citizen_id` varchar(36) NOT NULL,
    `gender` bool NOT NULL,
    `dob` varchar(12),
    `address` varchar(300),
    `ward` varchar(70),
    `district` varchar(70),
    `city` varchar(70),
    `current_work_place` text NOT NULL,
    `education_level` text,
    `experience` text,
    `certificate` text,
    `google_drive_url` varchar(300),
    `rate` double,
    `slogan` text,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_citizen_id` (`citizen_id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `nursing`;
-- +goose StatementEnd
