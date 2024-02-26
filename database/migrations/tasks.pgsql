CREATE TABLE tasks(
    id SERIAL NOT NULL, 
    name VARCHAR(30),
    user_id int,
    assigner_id int, 
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP, 
    description text, 
    project_id int,   
    current_status status NOT NULL,
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

select tablename,indexname,tablespace,indexdef  from pg_indexes where tablename = 'tasks';