-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_notes_title ON notes.note (title); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_notes_title;
-- +goose StatementEnd
