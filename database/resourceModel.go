package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetResources(businessID string) (resourceOutput []types.Resource, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.resource WHERE business_id = $1;", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	resourceOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Resource])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetResource(id string) (resourceOutput types.Resource, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.resource WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	resourceOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Resource])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteResource(id string) (err error) {

	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_public.resource WHERE id = $1;", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateResource(resourceInput types.Resource) (resourceId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`INSERT INTO risky_public.resource(name, description, cost, unit, total, resource_type, business_id)  values($1, $2, $3, $4, $5, $6, $7) RETURNING id;`,
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

func (m *DBManager) UpdateResource(resourceInput types.Resource) (err error) {

	tx, err := m.DBPool.Begin(context.Background())

	if err != nil {
		log.Println(err)
		return
	}

	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(),
		`UPDATE risky_public.resource SET name = $2, description = $3, cost = $4, unit = $5, total = $6, resource_type = $7, business_id = $8 WHERE id = $1;`,
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

	//Write a sql transaction to upadate all liability entries using this resource by multiplying the resource cost by the number of resources used in the liability
	_, err = tx.Exec(context.Background(),
		`UPDATE risky_public.liability SET cost = resource_quantity * $2 WHERE resource_id = $1;`,
		resourceInput.ID,
		resourceInput.Cost)
	if err != nil {
		log.Println(err)
		return
	}

	err = tx.Commit(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	return
}
