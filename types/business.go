package types

import (
	"time"

	"github.com/google/uuid"
)

type Business struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Revenue   float32   `json:"revenue"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Businesses []Business
