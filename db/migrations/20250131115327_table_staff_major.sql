-- +goose Up
-- +goose StatementBegin
CREATE TABLE `majors_staffs` (
    `major_id` varchar(36) NOT NULL,
    `staff_id` varchar(36) NOT NULL,
    PRIMARY KEY (`major_id`, `staff_id`),
    CONSTRAINT `fk_major_nurse_major` FOREIGN KEY (`major_id`) 
        REFERENCES `majors` (`id`) 
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `fk_major_nurse_nurse` FOREIGN KEY (`staff_id`) 
        REFERENCES `nurses` (`id`) 
        ON UPDATE CASCADE ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `majors_staffs`;
-- +goose StatementEnd
