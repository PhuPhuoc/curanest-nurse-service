-- +goose Up
-- +goose StatementBegin
CREATE TABLE `majors` (
    `id` varchar(36) NOT NULL,
    `name` varchar(300) NOT NULL,
    PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `majors`;
-- +goose StatementEnd
