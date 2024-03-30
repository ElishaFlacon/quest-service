-- MIGRATION FROM 30.03.2024





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
	value VARCHAR(255) NOT NULL,
	description VARCHAR(255),
	role VARCHAR(255) DEFAULT 'all',
	visible BOOLEAN DEFAULT true,
	PRIMARY KEY (id_indicator)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest_indicator (
	id_quest_indicator SERIAL NOT NULL UNIQUE,
	id_quest INT NOT NULL,
	id_indicator INT NOT NULL,
	PRIMARY KEY (id_quest_indicator)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS result (
	id_result SERIAL NOT NULL UNIQUE,
    id_indicator INT NOT NULL,
	id_launch_quest INT NOT NULL,
	id_from_user INT NOT NULL,
	id_to_user INT NOT NULL,
	value VARCHAR(255) NOT NULL,
	PRIMARY KEY (id_result)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest (
	id_quest SERIAL NOT NULL UNIQUE,
	available BOOLEAN DEFAULT true,
	name VARCHAR(255) NOT NULL,
	description VARCHAR(255),
	PRIMARY KEY (id_quest)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS launch_quest (
	id_launch_quest SERIAL NOT NULL UNIQUE,
	id_quest INT NOT NULL,
	id_team INT NOT NULL,
	available BOOLEAN DEFAULT true,
	start_at INT NOT NULL,
	end_at INT NOT NULL,
	PRIMARY KEY (id_launch_quest)
);
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE quest_indicator ADD CONSTRAINT quest_indicator_fk1 FOREIGN KEY (id_quest) REFERENCES quest(id_quest);
ALTER TABLE quest_indicator ADD CONSTRAINT quest_indicator_fk2 FOREIGN KEY (id_indicator) REFERENCES indicator(id_indicator);
ALTER TABLE result ADD CONSTRAINT result_fk1 FOREIGN KEY (id_launch_quest) REFERENCES launch_quest(id_launch_quest);
ALTER TABLE launch_quest ADD CONSTRAINT launch_quest_fk1 FOREIGN KEY (id_quest) REFERENCES quest(id_quest);
-- +goose StatementEnd





-- +goose Down

-- +goose StatementBegin
ALTER TABLE quest_indicator DROP CONSTRAINT IF EXISTS quest_indicator_fk1;
ALTER TABLE quest_indicator DROP CONSTRAINT IF EXISTS quest_indicator_fk2;
ALTER TABLE result DROP CONSTRAINT IF EXISTS result_fk1;
ALTER TABLE launch_quest DROP CONSTRAINT IF EXISTS launch_quest_fk1;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS example;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS indicator;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS quest_indicator;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS result;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS quest;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS launch_quest;
-- +goose StatementEnd

