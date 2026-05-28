CREATE TYPE challenge_level AS ENUM ('light', 'medium', 'hard');

CREATE TABLE workspace_rewards (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    workspace_id  UUID NOT NULL UNIQUE REFERENCES workspaces(id) ON DELETE CASCADE,
    points_light  INT  NOT NULL DEFAULT 10,
    points_medium INT  NOT NULL DEFAULT 25,
    points_hard   INT  NOT NULL DEFAULT 50,
    prize_light   TEXT,
    prize_medium  TEXT,
    prize_hard    TEXT
);