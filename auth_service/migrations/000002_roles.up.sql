CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO roles (name) VALUES ('System Administrator');
INSERT INTO roles (name) VALUES ('HR Manager');
INSERT INTO roles (name) VALUES ('Recruiter');
INSERT INTO roles (name) VALUES ('Department Manager');
INSERT INTO roles (name) VALUES ('Employee');
INSERT INTO roles (name) VALUES ('Accountant');
