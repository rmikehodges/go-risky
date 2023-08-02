package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetLiabilities(businessID string) (liabilityOutput []types.Liability, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.liability WHERE business_id = $1;)", businessID)
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

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.liability WHERE business_id = $1 AND impact_id = $2;", businessID, impactId)
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

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.liability WHERE business_id = $1 AND threat_id = $2;", businessID, threatId)
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

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.liability WHERE business_id = $1 AND mitigation_id = $2;", businessID, mitigationId)
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

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.liability WHERE id = $1;", id)
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

	tx, err := m.DBPool.Begin(context.Background())

	if err != nil {
		log.Println(err)
		return
	}

	defer tx.Rollback(context.Background())

	var resourceCost float32
	err = tx.QueryRow(context.Background(),
		`SELECT cost FROM risky_public.resource WHERE id = $1;`, liabilityInput.ResourceID).Scan(resourceCost)

	if err != nil {
		log.Println(err)
		return
	}

	liabilityInput.Cost = *liabilityInput.ResourceQuantity * resourceCost

	err = tx.QueryRow(context.Background(),
		`INSERT INTO risky_public.liability(
			name, description, category, impact_type, resopurce_quantity, 
			cost, type, business_id, resource_id, 
			threat_id) 
		values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id ;`,
		liabilityInput.Name,
		liabilityInput.Description,
		liabilityInput.Category,
		liabilityInput.ImpactType,
		liabilityInput.ResourceQuantity,
		liabilityInput.Type,
		liabilityInput.BusinessID,
		liabilityInput.ResourceID,
		liabilityInput.ThreatID).Scan(&liabilityId)
	if err != nil {
		log.Println(err)
		return
	}

	err = tx.Commit(context.Background())
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
	tx, err := m.DBPool.Begin(context.Background())

	if err != nil {
		log.Println(err)
		return
	}

	defer tx.Rollback(context.Background())

	var resourceCost float32
	err = tx.QueryRow(context.Background(),
		`SELECT cost FROM risky_public.resource WHERE id = $1;`, liabilityInput.ResourceID).Scan(resourceCost)

	if err != nil {
		log.Println(err)
		return
	}

	liabilityInput.Cost = *liabilityInput.ResourceQuantity * resourceCost

	_, err = tx.Exec(context.Background(),
		`UPDATE  risky_public.liability
		SET name = $1, description = $2, category = $3, impact_type = $4, resopurce_quantity = $5,
			cost = $6, type = $7, business_id = $8, resource_id = $9, threat_id = $10
			WHERE id = $11;;`,
		liabilityInput.Name,
		liabilityInput.Description,
		liabilityInput.Category,
		liabilityInput.ImpactType,
		liabilityInput.ResourceQuantity,
		liabilityInput.Type,
		liabilityInput.BusinessID,
		liabilityInput.ResourceID,
		liabilityInput.ThreatID,
		liabilityInput.ID)
	if err != nil {
		log.Println(err)
		return
	}

	err = tx.Commit(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	return
}
