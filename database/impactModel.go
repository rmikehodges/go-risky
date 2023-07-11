package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetImpacts(businessID string) (impactOutput []types.Impact, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, threat_id, exploitation_cost, mitigation_cost, created_at FROM risky_public.impacts(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	impactOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Impact])
	if err != nil {
		log.Printf("GetImpacts Error %s:", err)
		return
	}

	return
}

func (m *DBManager) GetImpact(id string) (impactOutput types.Impact, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, threat_id, exploitation_cost, mitigation_cost, created_at FROM risky_public.get_impact(fn_impact_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	impactOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Impact])
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

func (m *DBManager) CreateImpact(impactInput types.Impact) (impactId string, err error) {
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

func (m *DBManager) UpdateImpact(impactInput types.Impact) (err error) {

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
