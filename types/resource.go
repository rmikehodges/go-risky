package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type Resource struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  zeronull.Text `json:"description"`
	Cost         float32       `json:"cost" db:"cost"`
	Unit         string        `json:"unit" db:"unit"`
	Total        float32       `json:"total"`
	ResourceType string        `json:"resourceType" db:"resource_type"`
	BusinessID   uuid.UUID     `json:"businessId" db:"business_id"`
	CreatedAt    time.Time     `json:"createdAt" db:"created_at"`
}

type Resources []Resource
