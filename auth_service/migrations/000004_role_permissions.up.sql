CREATE TABLE role_permissions (
    role_id INT NOT NULL,
    permission_id INT NOT NULL,
    PRIMARY KEY (role_id, permission_id),
    FOREIGN KEY (role_id) REFERENCES roles(id),
    FOREIGN KEY (permission_id) REFERENCES permissions(id)
);

INSERT INTO permissions (name) VALUES ('CREATE_USER');
INSERT INTO permissions (name) VALUES ('VIEW_USER');
INSERT INTO permissions (name) VALUES ('UPDATE_USER');
INSERT INTO permissions (name) VALUES ('DELETE_USER');
INSERT INTO permissions (name) VALUES ('VIEW_PAYROLL');
INSERT INTO permissions (name) VALUES ('MANAGE_PAYROLL');
