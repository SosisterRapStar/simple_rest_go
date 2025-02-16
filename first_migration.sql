CREATE SCHEMA notes;


CREATE TABLE notes.note (
    id UUID PRIMARY KEY,
    title VARCHAR(50),
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
)
