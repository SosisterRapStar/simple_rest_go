-- +goose Up
-- +goose StatementBegin
CREATE INDEX time_id ON notes.note (created_at, id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX time_id;
-- +goose StatementEnd
