-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA notes;

CREATE TABLE notes.note (
    id UUID PRIMARY KEY,
    title VARCHAR(50),
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN 
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON notes.note
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER set_timestamp ON notes.note;
DROP FUNCTION trigger_set_timestamp();
DROP TABLE notes.note;
DROP SCHEMA notes;
-- +goose StatementEnd