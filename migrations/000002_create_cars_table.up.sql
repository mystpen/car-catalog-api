CREATE TABLE IF NOT EXISTS cars(
    car_id bigserial PRIMARY KEY,
    regNum string NOT NULL,
    mark string NOT NULL,
    model string NOT NULL,
    year integer,
    owner_id bigint,
    FOREIGN KEY (owner_id) REFERENCES people (person_id) ON DELETE SET NULL
);