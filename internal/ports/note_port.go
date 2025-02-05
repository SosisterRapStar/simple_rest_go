package ports

import "first-proj/application/domain"

// internal port
type NoteUseCase interface {
	CreateNewNote(title string, description string, createdAt string) *domain.Note
}
