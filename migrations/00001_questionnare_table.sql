-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS example (
 id SERIAL PRIMARY KEY,
 value VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS example;
-- +goose StatementEnd
