package notes

var (
	createNoteQuery = `
	INSERT INTO notes.note n (
		name, 
		content,
		expires_at
	)
	VALUES (
		$1,
		$2, 
		$3
	)
	RETURNING
	n.id,
	n.name,
	n.content,
	n.expires_at
	`

	deleteNoteQuery = `
	UPDATE notes.note n 
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE n.ID = $1;

	UPDATE notes.note_tag nt 
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE nt.note_id = $1; 
	`

	getUserNotesQuery = ` 
	SELECT
	name
	id
	, name
	, content
	, expires_at
	, created_at
	FROM notes.note n 
	WHERE n.user_id = $1 
	AND n.deleted_at is NULL;
	`

	getNoteTags = `
	SELECT
		t.name 
	FROM notes.tag t 
	JOIN notes.note_tag nt 
	ON t.id = nt.tag_id
	WHERE 
		nt.deleted_at is NULL 
	AND
		nt.note_id = $1
	`
	updateQueryNote = `
	UPDATE notes.note n 
	SET 
	updated_at = CURRENT_TIMESTAMP
	, content = COALESCE($2, n.content)
	, name = COALESCE($3, n.name)
	, expires_at = COALESCE($4, n.expires_at)
	WHERE n.id = $1 
	RETURNING
	n.id,
	n.name,
	n.content,
	n.expires_at
	`

	deleteTagQuery = `
	UPDATE notes.note n 
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE n.ID = $1;

	UPDATE notes.note_tag nt 
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE nt.note_id = $1; 
	`
)
