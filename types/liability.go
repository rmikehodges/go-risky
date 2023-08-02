package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type Liability struct {
	ID               uuid.UUID     `json:"id"`
	Name             string        `json:"name"`
	Description      zeronull.Text `json:"description"`
	Category         string        `json:"category"`
	ImpactType       string        `json:"impactType" db:"impact_type"`
	ResourceQuantity *float32      `json:"quantity" db:"resource_quantity"`
	Cost             float32       `json:"cost"`
	Type             string        `json:"type"`
	BusinessID       uuid.UUID     `json:"businessId" db:"business_id"`
	ResourceID       *uuid.UUID    `json:"resourceId" db:"resource_id"`
	ThreatID         *uuid.UUID    `json:"threatId" db:"threat_id"`
	CreatedAt        time.Time     `json:"createdAt" db:"created_at"`
}

type Liabilities []Liability
