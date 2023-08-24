BEGIN;

CREATE TABLE IF NOT EXISTS city (
    name varchar(64) PRIMARY KEY NOT NULL
);

INSERT INTO city (name)
SELECT DISTINCT(city) 
FROM institution;

ALTER TABLE institution
    ADD CONSTRAINT fk_institution_city FOREIGN KEY (city) REFERENCES city (name);

ALTER TABLE city ADD COLUMN ts tsvector
GENERATED ALWAYS AS
    (setweight(to_tsvector('russian', coalesce(name, '')), 'A')) STORED;

CREATE INDEX city_gin_ids ON city USING GIN (ts);

COMMIT;