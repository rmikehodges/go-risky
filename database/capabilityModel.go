package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetCapabilities(businessID string) (capabilityOutput []types.Capability, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, created_at FROM risky_public.capabilities(fn_business_id => $1)", businessID)
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

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, created_at FROM risky_public.get_capability(fn_capability_id => $1)", id)
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

	_, err = m.DBPool.Exec(context.Background(), "select risky_public.delete_capability(fn_capability_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateCapability(capabilityInput types.Capability) (capabilityId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`select * FROM risky_public.create_capability(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3)`,
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
		`select risky_public.update_capability(
			fn_capability_id => $1,
			fn_name => $2, 
			fn_description => $3,
			fn_business_id => $4)`,
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
