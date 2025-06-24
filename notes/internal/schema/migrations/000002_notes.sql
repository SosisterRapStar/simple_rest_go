-- +migration Up
CREATE SCHEMA IF NOT EXISTS notes;

CREATE TABLE IF NOT EXISTS notes.note (
    id uuid NOT NULL DEFAULT uuid_generate_v7() PRIMARY KEY ,
    user_id UUID NOT NULL                                   ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    updated_at TIMESTAMP                                    ,
    deleted_at TIMESTAMP
);


-- +migration Down
