package database

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type LiabilityModel struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  zeronull.Text `json:"description"`
	Quantity     float32       `json:"quantity"`
	Cost         float32       `json:"cost"`
	BusinessID   uuid.UUID     `json:"businessId" db:"business_id"`
	MitigationID uuid.UUID     `json:"mitigationId" db:"mitigation_id"`
	ResourceID   uuid.UUID     `json:"resourceId" db:"resource_id"`
	ThreatID     uuid.UUID     `json:"threatId" db:"threat_id"`
	ImpactID     uuid.UUID     `json:"impactId" db:"impact_id"`
	CreatedAt    time.Time     `json:"createdAt" db:"created_at"`
}

func (m *DBManager) GetLiabilities(businessID string) (liabilityOutput []LiabilityModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, mitigation_id, resource_id, threat_id, impact_id, created_at FROM risky_public.liabilities(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	liabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[LiabilityModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetLiability(id string) (liabilityOutput LiabilityModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, mitigation_id, resource_id, threat_id, impact_id, created_at FROM risky_public.get_liability(fn_liability_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	liabilityOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[LiabilityModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetLiabilityByImpactId(impactId string) (liabilityOutput []LiabilityModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, mitigation_id, resource_id, threat_id, impact_id, created_at FROM risky_public.get_liability_by_impact_id(fn_impact_id => $1)", impactId)
	if err != nil {
		log.Println(err)
		return
	}

	liabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[LiabilityModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteLiability(id string) (err error) {

	_, err = m.DBPool.Query(context.Background(), "select risky_public.delete_liability(fn_liability_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateLiability(liabilityInput LiabilityModel) (err error) {

	_, err = m.DBPool.Query(context.Background(),
		`select risky_public.create_liability(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3, 
			fn_mitigation_id => $4, 
			fn_resource_id => $5, 
			fn_threat_id => $6, 
			fn_impact_id => $7)`,
		liabilityInput.Name,
		liabilityInput.Description,
		liabilityInput.BusinessID,
		liabilityInput.MitigationID,
		liabilityInput.ResourceID,
		liabilityInput.ThreatID,
		liabilityInput.ImpactID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateLiability(liabilityInput LiabilityModel) (err error) {

	_, err = m.DBPool.Query(context.Background(),
		`select risky_public.update_liability(
			fn_liability_id => $1
			fn_name => $2, 
			fn_description => $3, 
			fn_business_id => $4, 
			fn_mitigation_id => $5, 
			fn_resource_id => $6, 
			fn_threat_id => $7, 
			fn_impact_id => $8)`,
		liabilityInput.ID,
		liabilityInput.Name,
		liabilityInput.Description,
		liabilityInput.BusinessID,
		liabilityInput.MitigationID,
		liabilityInput.ResourceID,
		liabilityInput.ThreatID,
		liabilityInput.ImpactID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
