package database

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type BusinessModel struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Revenue   float32   `json:"revenue"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

func (m *DBManager) GetBusinesses() (businessOutput []BusinessModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, revenue, created_at FROM risky_public.businesses()")
	if err != nil {
		log.Println(err)
		return
	}

	businessOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[BusinessModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetBusiness(id string) (businessOutput BusinessModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, revenue, created_at FROM risky_public.get_business(fn_business_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	businessOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[BusinessModel])
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

func (m *DBManager) CreateBusiness(businessInput BusinessModel) (businessId string, err error) {

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

func (m *DBManager) UpdateBusiness(businessInput BusinessModel) (err error) {

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
