-- +goose Up
-- +goose StatementBegin
CREATE TABLE `nursing_service` (
    `nursing_id` varchar(36) NOT NULL,
    `service_id` varchar(36) NOT NULL,
    PRIMARY KEY (`nursing_id`, `service_id`),
    CONSTRAINT `fk_nursing_service_nursing` FOREIGN KEY (`nursing_id`) 
        REFERENCES `nursing` (`id`) 
        ON UPDATE CASCADE ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `nursing_service`;
-- +goose StatementEnd
