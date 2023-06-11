package database

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type CapabilityModel struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at"`
}


func (m *DBManager) GetCapabilities(businessID string) (capabilityOutput []CapabilityModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, created_at FROM risky_public.capabilities(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	capabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[CapabilityModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetCapability(id string) (capabilityOutput CapabilityModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, created_at FROM risky_public.get_capability(fn_capability_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	capabilityOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[CapabilityModel])
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

func (m *DBManager) CreateCapability(capabilityInput CapabilityModel) (capabilityId string, err error) {

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

func (m *DBManager) UpdateCapability(capabilityInput CapabilityModel) (err error) {

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
