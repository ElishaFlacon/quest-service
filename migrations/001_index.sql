-- Мне стало лень вести миграции, так что актуальная ифнормация о базе будет тут))
-- FULL DATABASE CODE IN ONE FILE
-- LAST UPDATE: 21.04.2024





-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS category (
    id_category SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS indicator (
    id_indicator SERIAL PRIMARY KEY,
    id_category INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    from_role VARCHAR(255),
    to_role VARCHAR(255),
    visible BOOLEAN DEFAULT true
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
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    available BOOLEAN DEFAULT true
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest (
    id_quest SERIAL PRIMARY KEY,
    id_template INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    available BOOLEAN DEFAULT true,
    start_at INT NOT NULL,
    end_at INT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest_team (
    id_quest_team SERIAL PRIMARY KEY,
    id_team VARCHAR(255) NOT NULL,
    id_quest INT NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest_team_user (
    id_quest_team INT NOT NULL,
    id_user VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS result (
    id_result SERIAL PRIMARY KEY,
    id_indicator INT NOT NULL,
    id_quest INT NOT NULL,
    id_from_user VARCHAR(255) NOT NULL,
    id_to_user VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL
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
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO category 
(name) VALUES 
('программирование'), 
('дизайн'),
('смешное');

INSERT INTO indicator 
(id_category, name, description, from_role, to_role, visible) VALUES
(1, 'Уменее программировать', 'Как хорошо программирует?', '', '', true), 
(2, 'Уменее рисовать', 'Как хорошо рисует?', '', '', true), 
(2, 'Уменее моделирования', 'Как хорошо моделирует?', '', '', true), 
(3, 'Тест ролей 1', 'Тест', 'MEMBER', 'PROJECT_OFFICE', true), 
(3, 'Тест ролей 2', 'Тест', 'MEMBER', 'INITIATOR', true), 
(3, 'Тест visible', 'Тест', '', '', false), 
(3, 'Уменее веселиться', 'Как хорошо веселиться?', '', '', true);

INSERT INTO template 
(name, description, available) VALUES
('Шаблон 2022', 'Описание шаблона 2022', true), 
('Шаблон 2023', 'Описание шаблона 2023', true), 
('Шаблон 2024', 'Описание шаблона 2024', true), 
('Шаблон тест available', 'Описание шаблона', false);

INSERT INTO template_indicator 
(id_template, id_indicator) VALUES 
(1, 1), 
(1, 2), 
(1, 3), 
(1, 5), 
(2, 2), 
(2, 3), 
(2, 5), 
(3, 1), 
(3, 6), 
(4, 1), 
(4, 3), 
(4, 4), 
(4, 7);

INSERT INTO quest 
(id_template, name, description, available, start_at, end_at) VALUES 
(1, 'Базовый тест 1', 'Описание Тестовый тест 1', true, 1712245346, 1717515746), 
(2, 'Базовый тест 2', 'Описание Тестовый тест 2', true, 1712245346, 1717515746), 
(2, 'Скрытый тест', 'Описание скрытого теста', false, 1712245346, 1717515746), 
(1, 'Будущий тест', 'Описание будущего теста', false, 1736005346, 1751643746), 
(3, 'Прошедший тест', 'Описание прошедшего теста', false, 1704382946, 1709566946);

INSERT INTO quest_team 
(id_team, id_quest) VALUES 
('50', 1),
('56', 1),
('30', 2),
('31', 3),
('98', 1),
('98', 2),
('98', 3),
('98', 4),
('98', 5);

INSERT INTO quest_team_user 
(id_quest_team , id_user) VALUES 
(1, '500'),
(2, '500'),
(3, '500'),
(4, '500'),
(5, '500'),
(6, '500'),
(7, '500'),
(8, '500'),
(9, '500');

INSERT INTO result 
(id_indicator, id_quest, id_from_user, id_to_user, value) VALUES 
(1, 1, '500', '524', '5'),
(1, 2, '500', '524', '4'),
(1, 3, '500', '524', '3'),
(1, 5, '500', '524', '2'),
(1, 1, '500', '525', '2'),
(1, 2, '500', '525', '3'),
(1, 3, '500', '525', '2'),
(1, 5, '500', '525', '3'),
(1, 1, '501', '601', '6'),
(2, 1, '501', '601', '6'),
(3, 1, '501', '601', '6'),
(4, 1, '501', '601', '6'),
(5, 1, '501', '601', '6'),
(6, 1, '501', '601', '6'),
(7, 1, '501', '601', '6');
-- +goose StatementEnd