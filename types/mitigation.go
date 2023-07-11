package types

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type Mitigation struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
	Implemented bool          `json:"implemented"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at"`
}
type Mitigations []Mitigation
