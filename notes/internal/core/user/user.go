package user

import "github.com/google/uuid"

type User struct {
	UserId uuid.UUID `json:"user_id"`
}
