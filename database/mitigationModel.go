package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetMitigations(businessID string) (mitigationOutput []types.Mitigation, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.mitigation WHERE business_id = $1", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	mitigationOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Mitigation])
	if err != nil {
		log.Printf("GetMitigations Error: %s", err)
		return
	}

	return
}

func (m *DBManager) GetMitigation(id string) (mitigationOutput types.Mitigation, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.mitigation WHERE id = $1", id)
	if err != nil {
		log.Printf("GetMitigation database error: %s", err)
		return
	}

	mitigationOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Mitigation])
	if err != nil {
		log.Printf("GetMitigation parsing error: %s", err)
		return
	}

	return
}

func (m *DBManager) DeleteMitigation(id string) (err error) {

	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_public.mitigation WHERE id = $1", id)
	if err != nil {
		log.Printf("DeleteMitigation error: %s", err)
		return
	}

	return
}

func (m *DBManager) CreateMitigation(mitigationInput types.Mitigation) (mitigationId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`INSERT INTO risky_public.mitigation(name, description, business_id, action_id, implemented) values($1, $2, $3, $4) RETURNING id;`,
		mitigationInput.Name,
		mitigationInput.Description,
		mitigationInput.BusinessID,
		mitigationInput.Implemented).Scan(&mitigationId)
	if err != nil {
		log.Printf("CreateMitigation error: %s", err)
		return
	}

	return
}

func (m *DBManager) UpdateMitigation(mitigationInput types.Mitigation) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`UPDATE risky_public.mitigation SET name = $2, description = $3, business_id = $4, implemented = $5 WHERE id = $1;`,
		mitigationInput.ID,
		mitigationInput.Name,
		mitigationInput.Description,
		mitigationInput.BusinessID,
		mitigationInput.Implemented)
	if err != nil {
		log.Printf("UpdateMitigation error: %s", err)
		return
	}

	return
}
