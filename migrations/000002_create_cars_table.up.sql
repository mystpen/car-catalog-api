CREATE TABLE IF NOT EXISTS cars(
    car_id bigserial PRIMARY KEY,
    regNum text NOT NULL,
    mark text NOT NULL,
    model text NOT NULL,
    year integer,
    owner_id bigint,
    FOREIGN KEY (owner_id) REFERENCES people (person_id) ON DELETE SET NULL
);