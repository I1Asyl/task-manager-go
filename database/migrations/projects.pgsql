CREATE TYPE status AS ENUM('waiting', 'in progress', 'finished', 'reserved');

CREATE TABLE projects(
    id SERIAL NOT NULL,
    descripton text  NOT NULL,
    name VARCHAR(30) NOT NULL,
    current_status status NOT NULL,
    team_id int NOT NULL,
    PRIMARY KEY(id), 
    CONSTRAINT projects_user_fk
        FOREIGN KEY(team_id) REFERENCES teams(id)
);