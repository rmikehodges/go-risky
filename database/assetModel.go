package database

import (
	"context"
	"log"

	"go-risky/types"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetAssets(businessID string) (assetOutput []types.Asset, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.asset WHERE business_id = $1", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	assetOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Asset])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetAsset(id string) (assetOutput types.Asset, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.asset WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	assetOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Asset])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateAsset(assetInput types.Asset) (assetId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		` INSERT INTO risky_public.asset(
			name, 
			description, 
			business_id) 
			values($1, $2, $3) 
			RETURNING id`,
		assetInput.Name,
		assetInput.Description,
		assetInput.BusinessID).Scan(&assetId)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteAsset(id string) (err error) {

	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_public.asset WHERE id=$1", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateAsset(assetInput types.Asset) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`UPDATE risky_public.asset SET 
		name = $2, 
		description = $3, 
		business_id = $4 
		WHERE id = $1;`,
		assetInput.ID,
		assetInput.Name,
		assetInput.Description,
		assetInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
