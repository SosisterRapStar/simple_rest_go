package repository

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

	getNoteQuery = ` 
	SELECT 
		n.name 
		, n.content
		, n.expires_at
		, array_agg(ntag.name::TEXT)
	FROM notes.note_tag nt
	JOIN notes.note n 
	ON n.id = nt.note_id
	JOIN notes.tag ntag
	ON nt.tag_id = ntag.id
	WHERE n.user_id = $1 AND nt.deleted_at is NULL
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
