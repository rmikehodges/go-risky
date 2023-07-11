package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetThreats(businessID string) (threatOutput []types.Threat, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id,created_at FROM risky_public.threats(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	threatOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Threat])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetThreat(id string) (threatOutput types.Threat, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description,business_id, created_at FROM risky_public.get_threat(fn_threat_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	threatOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Threat])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteThreat(id string) (err error) {

	_, err = m.DBPool.Exec(context.Background(), "select risky_public.delete_threat(fn_threat_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateThreat(threatInput types.Threat) (threatId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`select * FROM risky_public.create_threat(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3)`,
		threatInput.Name,
		threatInput.Description,
		threatInput.BusinessID).Scan(&threatId)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateThreat(threatInput types.Threat) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`select risky_public.update_threat(
			fn_threat_id => $1,
			fn_name => $2, 
			fn_description => $3,
			fn_business_id => $4)`,
		threatInput.ID,
		threatInput.Name,
		threatInput.Description,
		threatInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
