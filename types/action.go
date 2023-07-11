package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type Action struct {
	ID              uuid.UUID     `json:"id"`
	Name            string        `json:"name" binding:"required"`
	Description     zeronull.Text `json:"description"`
	CapabilityID    *uuid.UUID    `json:"capabilityId" db:"capability_id"`
	VulnerabilityID *uuid.UUID    `json:"vulnerabilityId" db:"vulnerability_id"`
	BusinessID      uuid.UUID     `json:"businessId" inding:"required" db:"business_id"`
	Complexity      zeronull.Text `json:"complexity"`
	AssetID         *uuid.UUID    `json:"assetId" db:"asset_id"`
	CreatedAt       time.Time     `json:"createdAt" db:"created_at"`
}

type Actions []Action
