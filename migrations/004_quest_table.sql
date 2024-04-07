-- MIGRATION FROM 08.04.2024





-- +goose Up

-- +goose StatementBegin
ALTER TABLE IF EXISTS template ADD COLUMN IF NOT EXISTS description VARCHAR(255);
-- +goose StatementEnd





-- +goose Down

-- +goose StatementBegin
ALTER TABLE IF EXISTS template DROP COLUMN IF EXISTS description;
-- +goose StatementEnd
