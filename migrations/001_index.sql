-- Мне стало лень вести миграции, так что актуальная ифнормация о базе будет тут))
-- Last Update - 19.05.2024



-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS category (
    id_category SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS indicator (
    id_indicator SERIAL PRIMARY KEY,
    id_category INT NOT NULL,
    name VARCHAR(128) NOT NULL,
    description VARCHAR(256),
    answers VARCHAR(128)[],
    from_role VARCHAR(32),
    to_role VARCHAR(32),
    available BOOLEAN DEFAULT true
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS template_indicator (
    id_template_indicator SERIAL PRIMARY KEY,
    id_template INT NOT NULL,
    id_indicator INT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS template (
    id_template SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    description VARCHAR(256),
    available BOOLEAN DEFAULT true
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest (
    id_quest SERIAL PRIMARY KEY,
    id_template INT NOT NULL,
    name VARCHAR(128) NOT NULL,
    description VARCHAR(256),
    available BOOLEAN DEFAULT true,
    start_at INT NOT NULL,
    end_at INT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest_team (
    id_quest_team SERIAL PRIMARY KEY,
    id_team VARCHAR(64) NOT NULL,
    id_quest INT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest_team_user (
    id_quest_team INT NOT NULL,
    id_user VARCHAR(64) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS result (
    id_result SERIAL PRIMARY KEY,
    id_indicator INT NOT NULL,
    id_quest INT NOT NULL,
    id_from_user VARCHAR(64) NOT NULL,
    id_to_user VARCHAR(64) NOT NULL,
    value VARCHAR(128) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE indicator ADD CONSTRAINT category_fk1 FOREIGN KEY (id_category) REFERENCES category(id_category);
ALTER TABLE template_indicator ADD CONSTRAINT template_indicator_fk1 FOREIGN KEY (id_template) REFERENCES template(id_template);
ALTER TABLE template_indicator ADD CONSTRAINT template_indicator_fk2 FOREIGN KEY (id_indicator) REFERENCES indicator(id_indicator);
ALTER TABLE quest ADD CONSTRAINT quest_fk1 FOREIGN KEY (id_template) REFERENCES template(id_template);
ALTER TABLE quest_team ADD CONSTRAINT team_quest_fk1 FOREIGN KEY (id_quest) REFERENCES quest(id_quest);
ALTER TABLE quest_team_user ADD CONSTRAINT quest_team_user_fk1 FOREIGN KEY (id_quest_team) REFERENCES quest_team(id_quest_team);
ALTER TABLE result ADD CONSTRAINT result_fk1 FOREIGN KEY (id_quest) REFERENCES quest(id_quest);
ALTER TABLE result ADD CONSTRAINT result_fk2 FOREIGN KEY (id_indicator) REFERENCES indicator(id_indicator);
-- +goose StatementEnd
