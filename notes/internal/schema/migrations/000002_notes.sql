-- +migrate Up
-- change to uuidv7 in future 
CREATE SCHEMA IF NOT EXISTS notes;

CREATE TABLE IF NOT EXISTS notes.note (
    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id UUID NOT NULL,
    name TEXT NOT NULL,
    content TEXT,
    expires_at TIMESTAMP, 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS notes.tag (
    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT NOT NULL,
    CONSTRAINT cnst_notes_tags_unique_tags UNIQUE (name),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS notes.note_tag (
    note_id UUID NOT NULL REFERENCES notes.note(id),
    tag_id UUID NOT NULL REFERENCES notes.tag(id),    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (note_id, tag_id)
);

CREATE TABLE IF NOT EXISTS notes.edge (
    tail_vertex UUID NOT NULL REFERENCES notes.note(id),
    head_vertex UUID NOT NULL REFERENCES notes.note(id),
    label TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP,
    PRIMARY KEY (tail_vertex, head_vertex)
);


-- +migrate Down
DROP SCHEMA IF EXISTS notes CASCADE;