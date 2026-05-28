CREATE TYPE task_status AS ENUM ('pending', 'completed');

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    challenge_id UUID NOT NULL,
    assigned_to  UUID NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    is_daily BOOLEAN NOT NULL DEFAULT FALSE,
    status task_status NOT NULL DEFAULT 'pending',
    deadline TIMESTAMPTZ
);