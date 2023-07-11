package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetBusinesses() (businessOutput []types.Business, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, revenue, created_at FROM risky_public.businesses()")
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

	rows, err := m.DBPool.Query(context.Background(), "select id,name, revenue, created_at FROM risky_public.get_business(fn_business_id => $1)", id)
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

	_, err = m.DBPool.Exec(context.Background(), "select risky_public.delete_business(fn_business_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateBusiness(businessInput types.Business) (businessId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`select risky_public.create_business(
			fn_name => $1, 
			fn_revenue => $2)`,
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
		`select risky_public.update_business(
			fn_business_id => $1,
			fn_name => $2, 
			fn_revenue => $3)`,
		businessInput.ID,
		businessInput.Name,
		businessInput.Revenue)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
