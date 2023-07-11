package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type Liability struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  zeronull.Text `json:"description"`
	Quantity     *float32      `json:"quantity"`
	Cost         *float32      `json:"cost"`
	Type         string        `json:"type"`
	ResourceType string        `json:"resourceType" db:"resource_type"`
	BusinessID   uuid.UUID     `json:"businessId" db:"business_id"`
	MitigationID *uuid.UUID    `json:"mitigationId" db:"mitigation_id"`
	DetectionID  *uuid.UUID    `json:"detectionId" db:"detection_id"`
	ResourceID   *uuid.UUID    `json:"resourceId" db:"resource_id"`
	ImpactID     *uuid.UUID    `json:"impactId" db:"impact_id"`
	ThreatID     *uuid.UUID    `json:"threatId" db:"threat_id"`
	CreatedAt    time.Time     `json:"createdAt" db:"created_at"`
}

type Liabilities []Liability
