CREATE TABLE IF NOT EXISTS people(
    person_id bigserial PRIMARY KEY,
    name string NOT NULL,
    surname string NOT NULL,
    patronymic string
);