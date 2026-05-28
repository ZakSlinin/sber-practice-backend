CREATE TABLE user_achievements (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    achievement_id UUID NOT NULL,
    earned_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (user_id, achievement_id)
);