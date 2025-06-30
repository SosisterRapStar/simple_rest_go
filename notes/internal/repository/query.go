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
)
