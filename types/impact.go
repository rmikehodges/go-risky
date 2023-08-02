package types

import (
	"time"

	"github.com/google/uuid"
)

// make me a const aray in golang
var ImpactTypes = [2]string{"Exploitation", "Mitigation"}

type Impact struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	BusinessID  uuid.UUID `json:"businessId" db:"business_id"`
	ThreatID    uuid.UUID `json:"threatId" db:"threat_id"`
	ImpactType  string    `json:"impactType" db:"impact_type"`
	Cost        float32   `json:"cost"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}

type Impacts []Impact
