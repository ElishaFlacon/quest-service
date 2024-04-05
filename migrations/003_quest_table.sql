-- MIGRATION FROM 05.04.2024





-- +goose Up

-- +goose StatementBegin
ALTER TABLE IF EXISTS indicator ADD COLUMN IF NOT EXISTS id_category INT NOT NULL UNIQUE;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE IF EXISTS team ADD COLUMN IF NOT EXISTS id_team_quest SERIAL NOT NULL UNIQUE;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE IF EXISTS team DROP CONSTRAINT team_pkey;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE IF EXISTS team ADD PRIMARY KEY (id_team_quest);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS category (
    id_category SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id_category)
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE team ADD CONSTRAINT team_quest_fk1 FOREIGN KEY (id_quest) REFERENCES quest(id_quest);
ALTER TABLE indicator ADD CONSTRAINT category_fk1 FOREIGN KEY (id_category) REFERENCES category(id_category);
-- +goose StatementEnd





-- +goose Down

-- +goose StatementBegin
ALTER TABLE result DROP CONSTRAINT IF EXISTS team_quest_fk1;
ALTER TABLE indicator DROP CONSTRAINT IF EXISTS category_fk1;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS category;
-- +goose StatementEnd








-- лень дописывать down для миграции
