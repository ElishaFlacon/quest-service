-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS indicator (
	id_indicator serial NOT NULL UNIQUE,
	value varchar(255) NOT NULL,
	description varchar(255) NOT NULL,
	type varchar(255) NOT NULL,
	role varchar(255) NOT NULL,
	visible boolean NOT NULL,
	PRIMARY KEY (id_indicator)
);
-- +goose StatementEnd


-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest_indicator (
	id_quest_indicator serial NOT NULL UNIQUE,
	id_quest int NOT NULL,
	id_indicator int NOT NULL,
	PRIMARY KEY (id_quest_indicator)
);
-- +goose StatementEnd


-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS result (
	id_result serial NOT NULL UNIQUE,
    id_indicator int NOT NULL,
	id_launch_quest int NOT NULL,
	id_from_user int NOT NULL,
	id_to_user int NOT NULL,
	value varchar(255) NOT NULL,
	PRIMARY KEY (id_result)
);
-- +goose StatementEnd


-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest (
	id_quest serial NOT NULL UNIQUE,
	available boolean NOT NULL DEFAULT true,
	name varchar(255) NOT NULL,
	description varchar(255) NOT NULL,
	link varchar(255) NOT NULL,
	PRIMARY KEY (id_quest)
);
-- +goose StatementEnd


-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS launch_quest (
	id_launch_quest serial NOT NULL UNIQUE,
	id_quest int NOT NULL,
	id_team int NOT NULL,
	available boolean NOT NULL DEFAULT true,
	start_at int NOT NULL,
	end_at int NOT NULL,
	PRIMARY KEY (id_launch_quest)
);
-- +goose StatementEnd


-- +goose Up
-- +goose StatementBegin
ALTER TABLE quest_indicator ADD CONSTRAINT quest_indicator_fk1 FOREIGN KEY (id_quest) REFERENCES quest(id_quest);
ALTER TABLE quest_indicator ADD CONSTRAINT quest_indicator_fk2 FOREIGN KEY (id_indicator) REFERENCES indicator(id_indicator);
ALTER TABLE result ADD CONSTRAINT result_fk1 FOREIGN KEY (id_launch_quest) REFERENCES launch_quest(id_launch_quest);
ALTER TABLE launch_quest ADD CONSTRAINT launch_quest_fk1 FOREIGN KEY (id_quest) REFERENCES quest(id_quest);
-- +goose StatementEnd
