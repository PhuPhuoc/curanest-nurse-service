-- +goose Up
-- +goose StatementBegin
CREATE TABLE `nurses` (
    `id` varchar(36) NOT NULL,
    `major_id` varchar(36) NOT NULL,
    `citizen_id` varchar(36) NOT NULL,
    `gender` bool NOT NULL,
    `dob` varchar(12) DEFAULT NULL,
    `address` varchar(300) NOT NULL,
    `ward` varchar(70) NOT NULL,
    `district` varchar(70) NOT NULL,
    `city` varchar(70) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_citizen_id` (`citizen_id`),
    CONSTRAINT `fk_major_nurse` FOREIGN KEY (`major_id`) REFERENCES `majors` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `nurses`;
-- +goose StatementEnd
