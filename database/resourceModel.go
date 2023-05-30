package database

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type ResourceModel struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  zeronull.Text `json:"description"`
	Cost         float32       `json:"cost" db:"cost"`
	Unit         string        `json:"unit" db:"unit"`
	Total        float32       `json:"total"`
	ResourceType string        `json:"resourceType" db:"resource_type"`
	BusinessID   uuid.UUID     `json:"businessId" db:"business_id"`
	CreatedAt    time.Time     `json:"createdAt" db:"created_at"`
}

func (m *DBManager) GetResources(businessID string) (resourceOutput []ResourceModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, cost, unit, total, resource_type, business_id, created_at FROM risky_public.resources(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	resourceOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[ResourceModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetResource(id string) (resourceOutput ResourceModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, cost, unit, total, resource_type, business_id, created_at FROM risky_public.get_resource(fn_resource_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	resourceOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[ResourceModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteResource(id string) (err error) {

	_, err = m.DBPool.Exec(context.Background(), "select risky_public.delete_resource(fn_resource_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateResource(resourceInput ResourceModel) (resourceId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`select * FROM risky_public.create_resource(
			fn_name => $1, 
			fn_description => $2, 
			fn_cost => $3, 
			fn_unit => $4, 
			fn_total => $5,
			fn_resource_type => $6, 
			fn_business_id => $7)`,
		resourceInput.Name,
		resourceInput.Description,
		resourceInput.Cost,
		resourceInput.Unit,
		resourceInput.Total,
		resourceInput.ResourceType,
		resourceInput.BusinessID).Scan(&resourceId)
	if err != nil {
		log.Printf("CreateResource error: %s", err)
		return
	}

	return
}

func (m *DBManager) UpdateResource(resourceInput ResourceModel) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`select risky_public.update_resource(
			fn_resource_id => $1,
			fn_name => $2, 
			fn_description => $3, 
			fn_cost=> $4, 
			fn_unit => $5, 
			fn_total => $6, 
			fn_resource_type => $7,
			fn_business_id => $8)`,
		resourceInput.ID,
		resourceInput.Name,
		resourceInput.Description,
		resourceInput.Cost,
		resourceInput.Unit,
		resourceInput.Total,
		resourceInput.ResourceType,
		resourceInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
