CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(50) NOT NULL,
    surname VARCHAR(50) NOT NULL,
    phone VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(50) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    is_admin boolean NOT NULL
);

INSERT INTO users(username, name, surname, phone, password, email, is_admin) VALUES ('asus', 'yera', 'Yera', 'Altay', 'Qqwerty1!', 'altayerasyl@gmail.com', true);

INSERT INTO users(username, name, surname, phone, password, email, is_admin) VALUES ('asyl', 'yera', 'Yera', 'Altays', 'Qqwerty1!', 'altayyerasyl@gmail.com', false);

INSERT INTO users(username, name, surname, phone, password, email, is_admin) VALUES ('asylasus', 'yera', 'Yera', 'Altayss', 'Qqwerty1!', 'ltayyerasyl@gmail.com', false);
