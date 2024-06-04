CREATE TABLE users(
    email VARCHAR(128) PRIMARY KEY,
    password VARCHAR(512) NOT NULL
);

CREATE TABLE employees(
    email VARCHAR(64) PRIMARY KEY,
    date_of_birth DATE NOT NULL
);