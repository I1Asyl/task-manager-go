CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    phone VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    is_admin boolean NOT NULL
);
CREATE INDEX IF NOT EXISTS users_username ON users(username);

INSERT INTO users(username, name, surname, phone, password, email, is_admin) VALUES ('asus', 'yera', 'Yera', 'Altay', '73656372657414406d3848369e58541696a81a0dbe945227272a4ed57858385b585b8206ed8d', 'altayerasyl@gmail.com', true);

INSERT INTO users(username, name, surname, phone, password, email, is_admin) VALUES ('asyl', 'yera', 'Yera', 'Altays', '73656372657414406d3848369e58541696a81a0dbe945227272a4ed57858385b585b8206ed8d', 'altayyerasyl@gmail.com', false);

INSERT INTO users(username, name, surname, phone, password, email, is_admin) VALUES ('asylasus', 'yera', 'Yera', 'Altayss', '73656372657414406d3848369e58541696a81a0dbe945227272a4ed57858385b585b8206ed8d', 'ltayyerasyl@gmail.com', false);

CREATE TABLE IF NOT EXISTS teams(
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) UNIQUE NOT NULL, 
    description text NOT NULL

);

CREATE TABLE IF NOT EXISTS team_roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description text NOT NULL,
    can_edit_users boolean NOT NULL,
    can_edit_projects boolean NOT NULL, 
    can_edit_info boolean NOT NULL
);

INSERT INTO team_roles(name, description, can_edit_users, can_edit_projects, can_edit_info) VALUES ('manager', 'responsible for business logic of the software and how teams are organized', true, true, true);
INSERT INTO team_roles(name, description, can_edit_users, can_edit_projects, can_edit_info) VALUES ('junior', 'can view everything but can not edit anything', false, false, false);


CREATE TABLE IF NOT EXISTS users_teams(
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

CREATE INDEX IF NOT EXISTS users_teams_user_id ON users_teams(user_id); 
CREATE INDEX IF NOT EXISTS users_teams_user_id ON users_teams(team_id); 

CREATE TYPE status AS ENUM('waiting', 'in progress', 'finished', 'reserved');

CREATE TABLE IF NOT EXISTS projects(
    id SERIAL NOT NULL,
    description text  NOT NULL,
    name VARCHAR(30) UNIQUE NOT NULL,
    current_status status NOT NULL,
    team_id int NULL,
    PRIMARY KEY(id), 
    CONSTRAINT projects_user_fk
        FOREIGN KEY(team_id) REFERENCES teams(id) 
            ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS sessions(
    id SERIAL NOT NULL,
    first_token text  NOT NULL,
    token text NOT NULL,
    user_id int NOT NULL,
    PRIMARY KEY(id), 
    CONSTRAINT session_user_fk
        FOREIGN KEY(user_id) REFERENCES users(id)
            ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS sessions_first_token ON sessions(first_token);

CREATE TABLE IF NOT EXISTS tags(
    id SERIAL, 
    name VARCHAR(30),
    description text,
    PRIMARY KEY(id)
);

INSERT INTO tags(name, description) VALUES('dev', 'Development');

CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL NOT NULL, 
    name VARCHAR(30),
    user_id int,
    assigner_id int, 
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP, 
    description text, 
    project_id int,   
    current_status status NOT NULL,
    tag_id int,
    CONSTRAINT task_tag_fk 
        FOREIGN KEY(tag_id) REFERENCES tags(id)
            ON DELETE SET NULL,
    CONSTRAINT task_user_fk
        FOREIGN KEY(user_id) REFERENCES users(id)
            ON DELETE SET NULL, 
    CONSTRAINT task_assigner_fk
        FOREIGN KEY(assigner_id) REFERENCES users(id)
            ON DELETE SET NULL,     
    CONSTRAINT task_project_fk
        FOREIGN KEY(project_id) REFERENCES projects(id)
            ON DELETE SET NULL
);

