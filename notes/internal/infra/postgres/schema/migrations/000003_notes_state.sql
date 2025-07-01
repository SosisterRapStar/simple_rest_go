-- +migrate Up
ALTER TABLE notes.note ADD COLUMN IF NOT EXISTS state TEXT DEFAULT 'empty';
ALTER TABLE notes.note ADD CONSTRAINT IF NOT EXISTS cnst_notes_note_valid_state CHECK (state = ANY(ARRAY['empty'::text, 'done'::text, 'current'::text]));
-- +migrate Down
ALTER TABLE notes.note DROP COLUMN IF EXISTS state;