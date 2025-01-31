-- +goose Up
-- +goose StatementBegin
CREATE TABLE `nurse_profiles` (
    `id` varchar(36) NOT NULL,
    PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `nurse_profiles`;
-- +goose StatementEnd
