CREATE TABLE teams(
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) UNIQUE NOT NULL 
);

CREATE TABLE users_teams(
    user_id int NOT NULL, 
    team_id int NOT NULL,
    role_id int NOT NULL,
    PRIMARY KEY(user_id, team_id), 
    CONSTRAINT users_teams_user_fk
        FOREIGN KEY(user_id) REFERENCES users(id), 
    CONSTRAINT users_teams_team_fk
        FOREIGN KEY(team_id) REFERENCES teams(id),
    CONSTRAINT users_teams_role_fk
        FOREIGN KEY(role_id) REFERENCES team_roles(id)
);