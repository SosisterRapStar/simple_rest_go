-- +goose Up
-- +goose StatementBegin
ALTER TABLE notes.note ALTER COLUMN title TYPE TEXT;
ALTER TABLE notes.note ALTER COLUMN title SET NOT NULL;
ALTER TABLE notes.note ADD CONSTRAINT forbidden_symbols CHECK (title !~ '[@#$]');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE notes.note ALTER COLUMN title TYPE VARCHAR(50);
ALTER TABLE notes.note ALTER COLUMN title DROP NOT NULL;
ALTER TABLE notes.note DROP CONSTRAINT forbidden_symbols;
-- +goose StatementEnd
