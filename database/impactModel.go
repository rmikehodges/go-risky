package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetImpacts(businessID string) (impactOutput []types.Impact, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.impact WHERE business_id = $1", businessID)
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

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.impact WHERE id = $1", id)
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

	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_public.impact WHERE id = $1", id)
	if err != nil {
		log.Printf("DeleteImpact Error: %s", err)
		return
	}

	return
}

func (m *DBManager) CreateImpact(impactInput types.Impact) (impactId string, err error) {

	tx, err := m.DBPool.Begin(context.Background())


	if err != nil {
		log.Printf("CreateImpact error: %s", err)
		return
	}

	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(), `SELECT SUM(cost) FROM risky_public.liability WHERE business_id = $1 AND threat_id = $2 AND impact_type = $3`,
		impactInput.BusinessID, impactInput.ThreatID, impactInput.ImpactType).Scan(&impactInput.Cost)

	if err != nil {
		log.Printf("CreateImpact error: %s", err)
		return
	}

	err = tx.QueryRow(context.Background(),
		`INSERT INTO risky_public.impact
		(name,
		description, 
		business_id, 
		threat_id,
		impact_type, 
		cost) 
		values($1, $2, $3, $4, $5) RETURNING id;`,
		impactInput.Name,
		impactInput.Description,
		impactInput.BusinessID,
		impactInput.ThreatID,
		impactInput.ImpactType,
		impactInput.Cost).Scan(&impactId)
	if err != nil {
		log.Printf("CreateImpact error: %s", err)
		return
	}

	err = tx.Commit(context.Background())

	if err != nil {
		log.Printf("CreateImpact error: %s", err)
		return
	}

	return
}

func (m *DBManager) UpdateImpact(impactInput types.Impact) (err error) {
	tx, err := m.DBPool.Begin(context.Background())


	if err != nil {
		log.Printf("CreateImpact error: %s", err)
		return
	}

	defer tx.Rollback(context.Background())

	err = tx.QueryRow(context.Background(), `SELECT SUM(cost) FROM risky_public.liability WHERE business_id = $1 AND threat_id = $2 AND impact_type = $3`,
		impactInput.BusinessID, impactInput.ThreatID, impactInput.ImpactType).Scan(&impactInput.Cost)

	if err != nil {
		log.Printf("CreateImpact error: %s", err)
		return
	}

	_, err = tx.Exec(context.Background(),
		`UPDATE risky_public.impact 
		SET name = $2, 
		description = $3, 
		business_id = $4, 
		threat_id = $5, 
		impact_type = $6,
		cost = $7.  
		WHERE id = $1;`,
		impactInput.ID,
		impactInput.Name,
		impactInput.Description,
		impactInput.BusinessID,
		impactInput.ThreatID,
		impactInput.ImpactType,
		impactInput.Cost)
	if err != nil {
		log.Printf("CreateImpact error: %s", err)
		return
	}

	err = tx.Commit(context.Background())

	if err != nil {
		log.Printf("CreateImpact error: %s", err)
		return
	}

	return
}
