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

func (m *DBManager) GetLiabilities(businessID string) (liabilityOutput []LiabilityModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, quantity, type, resource_type, cost, business_id, mitigation_id, detection_id, resource_id, threat_id, impact_id, created_at FROM risky_public.liabilities(fn_business_id => $1)", businessID)
	if err != nil {
		log.Printf("GetLiabilities query error: %s", err)
		return
	}

	liabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[LiabilityModel])
	if err != nil {
		log.Printf("GetLiabilities CollectRows error: %s", err)
		return
	}

	return
}

func (m *DBManager) GetLiabilitiesByImpactId(businessID string, impactId string) (liabilityOutput []LiabilityModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, quantity, type, resource_type, cost, business_id, mitigation_id, detection_id, resource_id, threat_id, impact_id, created_at FROM risky_public.liabilities_by_impactId(fn_business_id => $1, fn_impact_id => $2)", businessID, impactId)
	if err != nil {
		log.Printf("GetLiabilitiesByImpactId Query error: %s", err)
		return
	}

	liabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[LiabilityModel])
	if err != nil {
		log.Printf("GetLiabilitiesByImpactId CollectRows error: %s", err)
		return
	}

	return
}

func (m *DBManager) GetLiabilitiesByThreatId(businessID string, threatId string) (liabilityOutput []LiabilityModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, quantity, type, resource_type, cost, business_id, mitigation_id, detection_id, resource_id, threat_id, impact_id, created_at FROM risky_public.liabilities_by_threat_Id(fn_business_id => $1, fn_threat_id => $2)", businessID, threatId)
	if err != nil {
		log.Printf("GetLiabilitiesByThreatId Query error: %s", err)
		return
	}

	liabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[LiabilityModel])
	if err != nil {
		log.Printf("GetLiabilitiesByThreatId CollectRows error: %s", err)
		return
	}

	return
}

func (m *DBManager) GetLiabilitiesByMitigationId(businessID string, mitigationId string) (liabilityOutput []LiabilityModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, quantity, type, resource_type, cost, business_id, mitigation_id, detection_id, resource_id, threat_id, impact_id, created_at FROM risky_public.liabilities_by_threatId(fn_business_id => $1, fn_mitigtation_id => $2)", businessID, mitigationId)
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

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, quantity, type, resource_type, cost, business_id, mitigation_id, detection_id, resource_id, threat_id, impact_id, created_at FROM risky_public.get_liability(fn_liability_id => $1)", id)
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

func (m *DBManager) CreateLiability(liabilityInput LiabilityModel) (liabilityId string, err error) {

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

func (m *DBManager) UpdateLiability(liabilityInput LiabilityModel) (err error) {

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
