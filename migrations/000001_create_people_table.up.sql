CREATE TABLE IF NOT EXISTS people(
    person_id bigserial PRIMARY KEY,
    name text NOT NULL,
    surname text NOT NULL,
    patronymic text
);