-- +goose Up
-- +goose StatementBegin
CREATE TABLE `staff_category` (
    `staff_id` varchar(36) NOT NULL,
    `category_id` varchar(36) NOT NULL,
    PRIMARY KEY (`category_id`, `staff_id`),
    CONSTRAINT `fk_staff_category_nurse` FOREIGN KEY (`staff_id`) 
        REFERENCES `nursing` (`id`) 
        ON UPDATE CASCADE ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `staff_category`;
-- +goose StatementEnd
