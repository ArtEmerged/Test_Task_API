CREATE TABLE people
(
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    patronymic VARCHAR(50),
    age SMALLSERIAL,
    gender VARCHAR(6)
);

CREATE TABLE nationalize (
    id SERIAL NOT NULL UNIQUE,
    person_id INT REFERENCES people(id) ON DELETE CASCADE NOT NULL,
    country_id VARCHAR(2) NOT NULL
);