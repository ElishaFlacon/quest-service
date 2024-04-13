-- MIGRATION FROM 14.04.2024





-- +goose Up

-- +goose StatementBegin
ALTER TABLE IF EXISTS indicator ADD COLUMN IF NOT EXISTS from_role VARCHAR(255);
ALTER TABLE IF EXISTS indicator ADD COLUMN IF NOT EXISTS to_role VARCHAR(255);
ALTER TABLE IF EXISTS indicator DROP COLUMN IF EXISTS role;
-- +goose StatementEnd





-- +goose Down

-- +goose StatementBegin
ALTER TABLE IF EXISTS indicator DROP COLUMN IF EXISTS from_role;
ALTER TABLE IF EXISTS indicator DROP COLUMN IF EXISTS to_role;
ALTER TABLE IF EXISTS indicator ADD COLUMN IF NOT EXISTS role VARCHAR(255);
-- +goose StatementEnd
