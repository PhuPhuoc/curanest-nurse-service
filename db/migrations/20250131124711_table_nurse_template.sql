-- +goose Up
-- +goose StatementBegin
CREATE TABLE `template` (
    `id` varchar(36) NOT NULL,
    PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `template`;
-- +goose StatementEnd
