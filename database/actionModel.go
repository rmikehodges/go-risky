package database

import (
	"context"
	"log"

	"go-risky/types"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetActions(businessID string) (actionOutput []types.Action, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, capability_id, vulnerability_id, business_id, complexity, asset_id, created_at FROM risky_public.actions(fn_business_id => $1)", businessID)
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

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, capability_id, vulnerability_id, business_id, complexity, asset_id, created_at FROM risky_public.get_action(fn_action_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	actionOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Action])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteAction(id string) (err error) {
	_, err = m.DBPool.Exec(context.Background(), "select risky_public.delete_action(fn_action_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateAction(actionInput types.Action) (createdAction string, err error) {
	err = m.DBPool.QueryRow(context.Background(),
		`select * FROM risky_public.create_action(
			fn_name => $1, 
			fn_description => $2, 
			fn_capability_id => $3, 
			fn_vulnerability_id => $4, 
			fn_business_id => $5, 
			fn_complexity => $6, 
			fn_asset_id => $7)`,
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
		`select risky_public.update_action(
			fn_action_id => $1,
			fn_name => $2, 
			fn_description => $3, 
			fn_capability_id => $4, 
			fn_vulnerability_id => $5, 
			fn_business_id => $6, 
			fn_complexity => $7, 
			fn_asset_id => $8)`,
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
