package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type AttackChain struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
	ThreatID    uuid.UUID     `json:"assetId" db:"threat_id"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at"`
}

type AttackChainOutputs []AttackChain
