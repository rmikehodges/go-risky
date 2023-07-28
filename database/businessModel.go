package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetBusinesses() (businessOutput []types.Business, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.business")
	if err != nil {
		log.Println(err)
		return
	}

	businessOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Business])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetBusiness(id string) (businessOutput types.Business, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.business WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	businessOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Business])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteBusiness(id string) (err error) {

	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_public.business WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateBusiness(businessInput types.Business) (businessId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`INSERT INTO risky_public.business(name, revenue)
		values($1, $2) RETURNING id;`,
		businessInput.Name,
		businessInput.Revenue).Scan(&businessId)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateBusiness(businessInput types.Business) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`UPDATE risky_public.business SET 
		name = $2, 
		revenue = $3 
		WHERE id = $1;`,
		businessInput.ID,
		businessInput.Name,
		businessInput.Revenue)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
