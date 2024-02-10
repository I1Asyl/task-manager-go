CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description text NOT NULL
);

INSERT INTO roles(name, description) VALUES ('manager', 'responsible for business logic of the software and how teams are organized');

DROP TABLE roles;