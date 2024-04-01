-- MIGRATION FROM 01.04.2024





-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS example (
    id SERIAL PRIMARY KEY,
    value VARCHAR(255)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS indicator (
    id_indicator SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    role VARCHAR(255) DEFAULT 'all',
    visible BOOLEAN DEFAULT true,
    PRIMARY KEY (id_indicator)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS template_indicator (
    id_template_indicator SERIAL NOT NULL UNIQUE,
    id_template INT NOT NULL,
    id_indicator INT NOT NULL,
    PRIMARY KEY (id_template_indicator)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS template (
    id_template SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    available BOOLEAN DEFAULT true,
    PRIMARY KEY (id_template)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest (
    id_quest SERIAL NOT NULL UNIQUE,
    id_template INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    available BOOLEAN DEFAULT true,
    start_at INT NOT NULL,
    end_at INT NOT NULL,
    PRIMARY KEY (id_quest)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS team (
    id_team INT NOT NULL,
    id_quest INT NOT NULL,
    PRIMARY KEY (id_team)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS result (
    id_result SERIAL NOT NULL UNIQUE,
    id_indicator INT NOT NULL,
    id_quest INT NOT NULL,
    id_from_user INT NOT NULL,
    id_to_user INT NOT NULL,
    value VARCHAR(255) NOT NULL,
    PRIMARY KEY (id_result)
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE template_indicator ADD CONSTRAINT template_indicator_fk1 FOREIGN KEY (id_template) REFERENCES template(id_template);
ALTER TABLE template_indicator ADD CONSTRAINT template_indicator_fk2 FOREIGN KEY (id_indicator) REFERENCES indicator(id_indicator);
ALTER TABLE quest ADD CONSTRAINT quest_fk1 FOREIGN KEY (id_template) REFERENCES template(id_template);
ALTER TABLE result ADD CONSTRAINT result_fk1 FOREIGN KEY (id_quest) REFERENCES quest(id_quest);
-- +goose StatementEnd





-- +goose Down

-- +goose StatementBegin
ALTER TABLE template_indicator DROP CONSTRAINT IF EXISTS template_indicator_fk1;
ALTER TABLE template_indicator DROP CONSTRAINT IF EXISTS template_indicator_fk2;
ALTER TABLE quest DROP CONSTRAINT IF EXISTS quest_fk1;
ALTER TABLE result DROP CONSTRAINT IF EXISTS result_fk1;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS example;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS indicator;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS template_indicator;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS template;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS quest;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS team;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS result;
-- +goose StatementEnd
