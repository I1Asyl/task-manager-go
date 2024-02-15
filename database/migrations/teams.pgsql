CREATE TABLE teams(
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) UNIQUE NOT NULL 
);

CREATE TABLE users_teams(
    id SERIAL PRIMARY KEY, 
    user_id int NOT NULL, 
    team_id int NOT NULL,
    CONSTRAINT users_teams_user_fk
        FOREIGN KEY(user_id) REFERENCES users(id), 
    CONSTRAINT users_teams_team_fk
        FOREIGN KEY(team_id) REFERENCES teams(id)        
);