CREATE TABLE sessions(
    id SERIAL NOT NULL,
    first_token text  NOT NULL,
    token text NOT NULL,
    user_id int NOT NULL,
    PRIMARY KEY(id), 
    CONSTRAINT session_user_fk
        FOREIGN KEY(user_id) REFERENCES users(id)
);