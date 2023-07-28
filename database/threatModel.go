package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetThreats(businessID string) (threatOutput []types.Threat, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.threat WHERE business_id = $1", businessID)
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

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.threat WHERE id = $1;", id)
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

	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_public.threat WHERE id = $1; ", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateThreat(threatInput types.Threat) (threatId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`INSERT INTO risky_public.threat(name, description, business_id) values($1, $2,$3) RETURNING id;`,
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
		`UPDATE risky_public.threat SET name = $2, description = $3,business_id = $4  WHERE id = $1;`,
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
