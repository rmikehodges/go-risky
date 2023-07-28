package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetCapabilities(businessID string) (capabilityOutput []types.Capability, err error) {

	rows, err := m.DBPool.Query(context.Background(), "    SELECT * FROM risky_public.capability WHERE business_id = $1;", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	capabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Capability])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetCapability(id string) (capabilityOutput types.Capability, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.capability WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	capabilityOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Capability])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteCapability(id string) (err error) {

	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_public.capability WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateCapability(capabilityInput types.Capability) (capabilityId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`INSERT INTO risky_public.capability(name, description, business_id) values($1, $2, $3) RETURNING id`,
		capabilityInput.Name,
		capabilityInput.Description,
		capabilityInput.BusinessID).Scan(&capabilityId)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateCapability(capabilityInput types.Capability) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`UPDATE risky_public.capability SET name = $2, description = $3, business_id = $4 WHERE id = $1`,
		capabilityInput.ID,
		capabilityInput.Name,
		capabilityInput.Description,
		capabilityInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
