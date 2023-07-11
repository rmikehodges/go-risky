package types

import (
	"time"

	"github.com/google/uuid"
)

type Impact struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Description      *string   `json:"description"`
	BusinessID       uuid.UUID `json:"businessId" db:"business_id"`
	ThreatID         uuid.UUID `json:"threatId" db:"threat_id"`
	ExploitationCost *float32  `json:"exploitationCost" db:"exploitation_cost"`
	MitigationCost   *float32  `json:"mitigationCost" db:"mitigation_cost"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
}

type Impacts []Impact
