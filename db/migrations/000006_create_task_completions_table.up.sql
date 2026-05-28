CREATE TABLE task_completions (
    id UUID PRIMARY KEY,
    task_id UUID NOT NULL UNIQUE,
    completed_by UUID NOT NULL,
    marked_by UUID NOT NULL,
    completed_at TIMESTAMPTZ NOT NULL
);