-- ENUMS
CREATE TYPE user_role AS ENUM ('admin', 'pm');
CREATE TYPE challenge_level AS ENUM ('light', 'medium', 'hard');
CREATE TYPE task_status AS ENUM ('pending', 'completed');

-- 1. workspaces (ни от чего не зависит)
CREATE TABLE workspaces (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE
);

-- 2. users (зависит от workspaces)
CREATE TABLE users (
    id            UUID      PRIMARY KEY DEFAULT gen_random_uuid(),
    workspace_id  UUID      NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    email         TEXT      NOT NULL,
    password_hash TEXT      NOT NULL,
    name          TEXT      NOT NULL,
    role          user_role NOT NULL DEFAULT 'pm',
    team          TEXT,
    total_points  INT       NOT NULL DEFAULT 0,

    UNIQUE (workspace_id, email)
);

-- 3. workspace_rewards (зависит от workspaces)
CREATE TABLE workspace_rewards (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workspace_id  UUID NOT NULL UNIQUE REFERENCES workspaces(id) ON DELETE CASCADE,
    points_light  INT  NOT NULL DEFAULT 10,
    points_medium INT  NOT NULL DEFAULT 25,
    points_hard   INT  NOT NULL DEFAULT 50,
    prize_light   TEXT,
    prize_medium  TEXT,
    prize_hard    TEXT
);

-- 4. challenges (зависит от workspaces и users)
CREATE TABLE challenges (
    id           UUID            PRIMARY KEY DEFAULT gen_random_uuid(),
    workspace_id UUID            NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
    title        TEXT            NOT NULL,
    description  TEXT,
    level        challenge_level NOT NULL,
    is_active    BOOLEAN         NOT NULL DEFAULT TRUE,
    created_by   UUID            NOT NULL REFERENCES users(id)
);

-- 5. tasks (зависит от challenges и users)
CREATE TABLE tasks (
    id           UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    challenge_id UUID        NOT NULL REFERENCES challenges(id) ON DELETE CASCADE,
    assigned_to  UUID        NOT NULL REFERENCES users(id),
    title        TEXT        NOT NULL,
    description  TEXT,
    is_daily     BOOLEAN     NOT NULL DEFAULT FALSE,
    status       task_status NOT NULL DEFAULT 'pending',
    deadline     TIMESTAMPTZ
);

-- 6. task_completions (зависит от tasks и users)
CREATE TABLE task_completions (
    id           UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    task_id      UUID        NOT NULL UNIQUE REFERENCES tasks(id) ON DELETE CASCADE,
    completed_by UUID        NOT NULL REFERENCES users(id),
    marked_by    UUID        NOT NULL REFERENCES users(id),
    completed_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 7. achievements (ни от чего не зависит)
CREATE TABLE achievements (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code        TEXT NOT NULL UNIQUE,
    title       TEXT NOT NULL,
    description TEXT,
    threshold   INT  NOT NULL DEFAULT 1
);

-- 8. user_achievements (зависит от users и achievements)
CREATE TABLE user_achievements (
    id             UUID        PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id        UUID        NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    achievement_id UUID        NOT NULL REFERENCES achievements(id),
    earned_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (user_id, achievement_id)
);