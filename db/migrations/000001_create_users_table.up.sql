CREATE TYPE user_role AS ENUM ('admin', 'pm');

CREATE TABLE users (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    workspace_id  UUID NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    email         TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    name          TEXT NOT NULL,
    role          user_role NOT NULL DEFAULT 'pm',
    team          TEXT,
    total_points  INT NOT NULL DEFAULT 0,

    UNIQUE (workspace_id, email)
);