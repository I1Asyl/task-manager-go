CREATE TABLE team_roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description text NOT NULL,
    can_edit_users boolean NOT NULL,
    can_edit_projects boolean NOT NULL
);

INSERT INTO team_roles(name, description, can_edit_users, can_edit_projects) VALUES ('manager', 'responsible for business logic of the software and how teams are organized', true, true);

