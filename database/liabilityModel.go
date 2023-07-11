package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetLiabilities(businessID string) (liabilityOutput []types.Liability, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, quantity, type, resource_type, cost, business_id, mitigation_id, detection_id, resource_id, threat_id, impact_id, created_at FROM risky_public.liabilities(fn_business_id => $1)", businessID)
	if err != nil {
		log.Printf("GetLiabilities query error: %s", err)
		return
	}

	liabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Liability])
	if err != nil {
		log.Printf("GetLiabilities CollectRows error: %s", err)
		return
	}

	return
}

func (m *DBManager) GetLiabilitiesByImpactId(businessID string, impactId string) (liabilityOutput []types.Liability, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, quantity, type, resource_type, cost, business_id, mitigation_id, detection_id, resource_id, threat_id, impact_id, created_at FROM risky_public.liabilities_by_impactId(fn_business_id => $1, fn_impact_id => $2)", businessID, impactId)
	if err != nil {
		log.Printf("GetLiabilitiesByImpactId Query error: %s", err)
		return
	}

	liabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Liability])
	if err != nil {
		log.Printf("GetLiabilitiesByImpactId CollectRows error: %s", err)
		return
	}

	return
}

func (m *DBManager) GetLiabilitiesByThreatId(businessID string, threatId string) (liabilityOutput []types.Liability, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, quantity, type, resource_type, cost, business_id, mitigation_id, detection_id, resource_id, threat_id, impact_id, created_at FROM risky_public.liabilities_by_threat_Id(fn_business_id => $1, fn_threat_id => $2)", businessID, threatId)
	if err != nil {
		log.Printf("GetLiabilitiesByThreatId Query error: %s", err)
		return
	}

	liabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Liability])
	if err != nil {
		log.Printf("GetLiabilitiesByThreatId CollectRows error: %s", err)
		return
	}

	return
}

func (m *DBManager) GetLiabilitiesByMitigationId(businessID string, mitigationId string) (liabilityOutput []types.Liability, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, quantity, type, resource_type, cost, business_id, mitigation_id, detection_id, resource_id, threat_id, impact_id, created_at FROM risky_public.liabilities_by_threatId(fn_business_id => $1, fn_mitigtation_id => $2)", businessID, mitigationId)
	if err != nil {
		log.Println(err)
		return
	}

	liabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Liability])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetLiability(id string) (liabilityOutput types.Liability, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, quantity, type, resource_type, cost, business_id, mitigation_id, detection_id, resource_id, threat_id, impact_id, created_at FROM risky_public.get_liability(fn_liability_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	liabilityOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Liability])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateLiability(liabilityInput types.Liability) (liabilityId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`select risky_public.create_liability(
			fn_name => $1, 
			fn_description => $2,
			fn_quantity => $3,
			fn_type => $4,
			fn_business_id => $5, 
			fn_mitigation_id => $6, 
			fn_resource_id => $7, 
			fn_detection_id => $8, 
			fn_threat_id => $10)
			fn_impact_id => $11)`,
		liabilityInput.Name,
		liabilityInput.Description,
		liabilityInput.Quantity,
		liabilityInput.Type,
		liabilityInput.BusinessID,
		liabilityInput.MitigationID,
		liabilityInput.ResourceID,
		liabilityInput.DetectionID,
		liabilityInput.ImpactID,
		liabilityInput.ThreatID).Scan(&liabilityId)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteLiability(id string) (err error) {

	_, err = m.DBPool.Exec(context.Background(), "select risky_public.delete_liability(fn_liability_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateLiability(liabilityInput types.Liability) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`select risky_public.update_liability(
			fn_liability_id => $1,
			fn_name => $2, 
			fn_description => $3,
			fn_quantity => $4,
			fn_type => $5,
			fn_business_id => $6, 
			fn_mitigation_id => $7, 
			fn_resource_id => $8, 
			fn_detection_id => $9, 
			fn_impact_id => $10, 
			fn_threat_id => $11)`,
		liabilityInput.ID,
		liabilityInput.Name,
		liabilityInput.Description,
		liabilityInput.Quantity,
		liabilityInput.Type,
		liabilityInput.BusinessID,
		liabilityInput.MitigationID,
		liabilityInput.ResourceID,
		liabilityInput.DetectionID,
		liabilityInput.ImpactID,
		liabilityInput.ThreatID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
