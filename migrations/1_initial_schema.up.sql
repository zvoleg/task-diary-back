CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS administration;
CREATE SCHEMA IF NOT EXISTS storage;
CREATE SCHEMA IF NOT EXISTS reference;

CREATE TABLE IF NOT EXISTS administration.users
(
    user_id       UUID PRIMARY KEY NOT NULL,
    account_id    UUID,
    name          VARCHAR(32),
    description   VARCHAR(512),
    role_id       UUID,
    created_at    TIMESTAMP NOT NULL,
    updated_at    TIMESTAMP,
    is_deleted    BOOLEAN
);

CREATE TABLE IF NOT EXISTS storage.tasks
(
    task_id       UUID PRIMARY KEY NOT NULL,
    board_id      UUID NOT NULL,
    title         VARCHAR(128),
    description   VARCHAR(512),
    type          VARCHAR(16),
    status        VARCHAR(16),
    author_id     UUID NOT NULL,
    created_at    TIMESTAMP NOT NULL,
    updated_at    TIMESTAMP,
    tags          VARCHAR(128),
    is_deleted    BOOLEAN
);

CREATE TABLE IF NOT EXISTS reference.permissions
(
    permission_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name          VARCHAR(32)
);

INSERT INTO reference.permissions (name) VALUES 
    ('user_read'),
    ('user_read_list'),
    ('user_write'),
    ('user_update'),
    ('user_delete'),
    ('role_read'),
    ('role_read_list'),
    ('board_read'),
    ('board_read_list'),
    ('board_write'),
    ('board_update'),
    ('board_delete'),
    ('task_read'),
    ('task_read_list'),
    ('task_write'),
    ('task_update'),
    ('task_delete');

CREATE TABLE IF NOT EXISTS reference.roles
(
    role_id       UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name          VARCHAR(32),
    description   VARCHAR(512)
);

INSERT INTO reference.roles (name, description) VALUES
    ('administrator', 'the role with full set of permission');

CREATE TABLE IF NOT EXISTS reference.role_permissions
(
    role_id       UUID NOT NULL,
    permission_id UUID NOT NULL,
    CONSTRAINT fk_role
        FOREIGN KEY(role_id)
            REFERENCES reference.roles(role_id),
    CONSTRAINT fk_permission_id
        FOREIGN KEY(permission_id)
            REFERENCES reference.permissions(permission_id)
);

DO $$DECLARE p UUID;
BEGIN
    FOR p IN SELECT permission_id from reference.permissions
    LOOP
        INSERT INTO reference.role_permissions (role_id, permission_id) VALUES
        ((SELECT role_id FROM reference.roles r WHERE r.name = 'administrator'), p);
    END LOOP;
END$$;