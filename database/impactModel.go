package database

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ImpactModel struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Description      *string   `json:"description"`
	BusinessID       uuid.UUID `json:"businessId" db:"business_id"`
	ThreatID         uuid.UUID `json:"threatId" db:"threat_id"`
	ExploitationCost *float32  `json:"exploitationCost" db:"exploitation_cost"`
	MitigationCost   *float32  `json:"mitigationCost" db:"mitigation_cost"`
	CreatedAt        time.Time `json:"createdAt" db:"created_at"`
}

func (m *DBManager) GetImpacts(businessID string) (impactOutput []ImpactModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, threat_id, exploitation_cost, mitigation_cost, created_at FROM risky_public.impacts(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	impactOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[ImpactModel])
	if err != nil {
		log.Printf("GetImpacts Error %s:", err)
		return
	}

	return
}

func (m *DBManager) GetImpact(id string) (impactOutput ImpactModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, threat_id, exploitation_cost, mitigation_cost, created_at FROM risky_public.get_impact(fn_impact_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	impactOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[ImpactModel])
	if err != nil {
		log.Printf("GetImpact error: %s", err)
		return
	}

	return
}

func (m *DBManager) DeleteImpact(id string) (err error) {

	_, err = m.DBPool.Exec(context.Background(), "select risky_public.delete_impact(fn_impact_id => $1)", id)
	if err != nil {
		log.Printf("DeleteImpact Error: %s", err)
		return
	}

	return
}

func (m *DBManager) CreateImpact(impactInput ImpactModel) (impactId string, err error) {
	//TODO: Add impact exploitation and mitigation cost

	err = m.DBPool.QueryRow(context.Background(),
		`select risky_public.create_impact(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3, 
			fn_threat_id => $4)`,
		impactInput.Name,
		impactInput.Description,
		impactInput.BusinessID,
		impactInput.ThreatID).Scan(&impactId)
	if err != nil {
		log.Printf("CreateImpact error: %s", err)
		return
	}

	return
}

func (m *DBManager) UpdateImpact(impactInput ImpactModel) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`select risky_public.update_impact(
			fn_impact_id => $1,
			fn_name => $2, 
			fn_description => $3, 
			fn_business_id => $4, 
			fn_threat_id => $5)`,
		impactInput.ID,
		impactInput.Name,
		impactInput.Description,
		impactInput.BusinessID,
		impactInput.ThreatID)
	if err != nil {
		log.Printf("UpdateImpact Error: %s", err)
		return
	}

	return
}
