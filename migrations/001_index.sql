-- Мне стало лень вести миграции, так что актуальная ифнормация о базе будет тут))
-- Last Update - 05.08.2024



-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS category (
    id_category SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    available BOOLEAN DEFAULT true
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
    name VARCHAR(128) NOT NULL,
    id_quest INT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest_team_user (
    id_quest_team INT NOT NULL,
    id_user VARCHAR(64) NOT NULL,
    name VARCHAR(128) NOT NULL,
    email VARCHAR(128) NOT NULL
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

-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_results_count_by_quest_id;
CREATE OR REPLACE FUNCTION get_results_count_by_quest_id(qid INT) 
	RETURNS INT AS $$
    BEGIN
        RETURN (SELECT COUNT(*) FROM (
            SELECT COUNT(DISTINCT "result".id_from_user)
            FROM "quest_team"
            INNER JOIN "quest_team_user" ON "quest_team".id_quest_team = "quest_team_user".id_quest_team
            INNER JOIN "result" ON "result".id_quest = qid
            WHERE "quest_team".id_quest = qid
            GROUP BY "result".id_from_user
        ) results_count);
    END; $$
    LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_users_count_by_quest_id;
CREATE OR REPLACE FUNCTION get_users_count_by_quest_id(qid INT) 
	RETURNS INT AS $$
    BEGIN
        RETURN (SELECT COUNT("quest_team_user".id_quest_team)
            FROM "quest_team"
            INNER JOIN "quest_team_user" ON "quest_team".id_quest_team = "quest_team_user".id_quest_team
            WHERE "quest_team".id_quest = qid
        );
    END; $$
    LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_quest_pass_percent;
CREATE OR REPLACE FUNCTION get_quest_pass_percent(qid INT) 
	RETURNS INT AS $$
	DECLARE res INT = (
		100.0 
			/ 
		NULLIF((SELECT get_users_count_by_quest_id(qid)), 0)
			* 
		NULLIF((SELECT get_results_count_by_quest_id(qid)), 0)
	); 
    BEGIN
		IF res IS NULL THEN
			RETURN 0;
		ELSE
			RETURN res;
        END IF;
    END; $$
    LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_results_count_by_team_id;
CREATE OR REPLACE FUNCTION get_results_count_by_team_id(tid VARCHAR(64), qid INT) 
	RETURNS INT AS $$
    BEGIN
        RETURN (SELECT COUNT(*) FROM (
            SELECT COUNT(DISTINCT "result".id_from_user)
            FROM "quest_team"
            INNER JOIN "quest_team_user" ON "quest_team".id_quest_team = "quest_team_user".id_quest_team
            INNER JOIN "result" ON "quest_team".id_quest = "result".id_quest AND "quest_team_user".id_user = "result".id_from_user
            WHERE "quest_team".id_team = tid AND "quest_team".id_quest = qid
            GROUP BY "result".id_from_user
        ) results_count);
    END; $$
    LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_users_count_by_team_and_quest_id;
CREATE OR REPLACE FUNCTION get_users_count_by_team_and_quest_id(tid VARCHAR(64), qid INT) 
	RETURNS INT AS $$
    BEGIN
        RETURN (
            SELECT COUNT("quest_team_user".id_quest_team)
            FROM "quest_team"
            INNER JOIN "quest_team_user" ON "quest_team".id_quest_team = "quest_team_user".id_quest_team
            WHERE  "quest_team".id_team = tid AND "quest_team".id_quest = qid
        );
    END; $$
    LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_team_pass_percent;
CREATE OR REPLACE FUNCTION get_team_pass_percent(tid VARCHAR(64), qid INT)
	RETURNS INT AS $$
	DECLARE res INT = (
		100.0
			/ 
		NULLIF((SELECT get_users_count_by_team_and_quest_id(tid, qid)), 0)
			* 
		NULLIF((SELECT get_results_count_by_team_id(tid, qid)), 0)
	); 
    BEGIN
		IF res IS NULL THEN
			RETURN 0;
		ELSE
			RETURN res;
        END IF;
    END; $$
    LANGUAGE plpgsql;
-- +goose StatementEnd

-- +goose StatementBegin
DROP FUNCTION IF EXISTS get_user_status;
CREATE OR REPLACE FUNCTION get_user_status(qid INT, uid VARCHAR(64)) 
	RETURNS BOOL AS $$
    BEGIN
        RETURN (
			SELECT
				CASE 
					WHEN COUNT("result".id_from_user) > 0
					THEN TRUE
					ELSE FALSE
				END AS status
			FROM "quest"
			INNER JOIN "quest_team" ON "quest".id_quest = "quest_team".id_quest
			INNER JOIN "result" ON "quest".id_quest = "result".id_quest
            WHERE "result".id_from_user = uid AND "quest".id_quest = qid 
        );
    END; $$
    LANGUAGE plpgsql;
-- +goose StatementEnd