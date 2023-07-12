package types

import (
	"time"

	"github.com/google/uuid"
)

type AttackChainStep struct {
	ID            uuid.UUID  `json:"id" db:"id"`
	BusinessID    uuid.UUID  `json:"businessId" db:"business_id"`
	ActionID      uuid.UUID  `json:"actionId" db:"action_id"`
	AssetID       *uuid.UUID `json:"assetId" db:"asset_id"`
	AttackChainID uuid.UUID  `json:"attackChainId" db:"attack_chain_id"`
	NextStep      *uuid.UUID `json:"nextStep" db:"next_step"`
	PreviousStep  *uuid.UUID `json:"previousStep" db:"previous_step"`
	DetectionID   uuid.UUID  `json:"detectionId" db:"detection_id"`
	MitigationID  *uuid.UUID `json:"mitigationId" db:"mitigation_id"`
	CreatedAt     time.Time  `json:"createdAt" db:"created_at"`
}

type AttackChainSteps []AttackChainStep
