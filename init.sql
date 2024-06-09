CREATE TABLE  IF NOT EXISTS users(
    email VARCHAR(64) PRIMARY KEY,
    password VARCHAR(512) NOT NULL
);

CREATE TABLE IF NOT EXISTS employees(
    email VARCHAR(64) PRIMARY KEY,
    date_of_birth TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS subs(
    id INTEGER PRIMARY KEY,
    user VARCHAR(64) REFERENCES users(email),
    employee VARCHAR(64) REFERENCES employees(email)
);