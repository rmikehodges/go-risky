package database

import (
	"context"
	"fmt"
	"log"

	"go-risky/types"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetActions(businessID string) (actionOutput []types.Action, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.action WHERE business_id = $1", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	actionOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Action])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetAction(id string) (actionOutput types.Action, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.action WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(rows)

	actionOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Action])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteAction(id string) (err error) {
	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_public.action WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateAction(actionInput types.Action) (createdAction string, err error) {
	err = m.DBPool.QueryRow(context.Background(),
		`INSERT INTO risky_public.action(
			name, 
			description, 
			capability_id, 
			vulnerability_id,
			 business_id, 
			 complexity, 
			 asset_id) 
			 values($1, $2, $3,$4, $5, $6, $7) 
			 RETURNING id`,
		actionInput.Name,
		actionInput.Description,
		actionInput.CapabilityID,
		actionInput.VulnerabilityID,
		actionInput.BusinessID,
		actionInput.Complexity,
		actionInput.AssetID).Scan(&createdAction)

	if err != nil {
		log.Printf("Error creating action %s", err)
		return
	}
	return
}

func (m *DBManager) UpdateAction(actionInput types.Action) (err error) {
	_, err = m.DBPool.Exec(context.Background(),
		`UPDATE risky_public.action SET 
		name = $2, 
		description = $3, 
		capability_id = $4, 
		vulnerability_id = $5, 
		business_id = $6, 
		complexity = $7, 
		asset_id = $8 
		WHERE id = $1`,
		actionInput.ID,
		actionInput.Name,
		actionInput.Description,
		actionInput.CapabilityID,
		actionInput.VulnerabilityID,
		actionInput.BusinessID,
		actionInput.Complexity,
		actionInput.AssetID)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
