CREATE TABLE IF NOT EXISTS todoes (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    text text NOT NULL,
    completed boolean NOT NULL DEFAULT false,
    tag text[] NOT NULL,
    version integer NOT NULL DEFAULT 1
    );