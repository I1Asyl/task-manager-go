CREATE TABLE teams(
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) UNIQUE NOT NULL, 
    description text NOT NULL

);

CREATE TABLE team_roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description text NOT NULL,
    can_edit_users boolean NOT NULL,
    can_edit_projects boolean NOT NULL, 
    can_edit_info boolean NOT NULL
);

INSERT INTO team_roles(name, description, can_edit_users, can_edit_projects, can_edit_info) VALUES ('manager', 'responsible for business logic of the software and how teams are organized', true, true, true);
INSERT INTO team_roles(name, description, can_edit_users, can_edit_projects, can_edit_info) VALUES ('junior', 'can view everything but can not edit anything', false, false, false);


CREATE TABLE users_teams(
    user_id int NOT NULL, 
    team_id int NOT NULL,
    role_id int NOT NULL DEFAULT 2,
    PRIMARY KEY(user_id, team_id), 
    CONSTRAINT users_teams_user_fk
        FOREIGN KEY(user_id) REFERENCES users(id) 
            ON DELETE CASCADE, 
    CONSTRAINT users_teams_team_fk
        FOREIGN KEY(team_id) REFERENCES teams(id)
            ON DELETE CASCADE, 
    CONSTRAINT users_teams_role_fk
        FOREIGN KEY(role_id) REFERENCES team_roles(id)
            ON DELETE SET DEFAULT
);