package database

import (
	"context"
	"log"

	"go-risky/types"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetAssets(businessID string) (assetOutput []types.Asset, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, created_at FROM risky_public.assets(fn_business_id => $1)", businessID)
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

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, created_at FROM risky_public.get_asset(fn_asset_id => $1)", id)
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
		`select * FROM risky_public.create_asset(
			fn_name => $1,
			fn_description => $2,
			fn_business_id => $3)`,
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

	_, err = m.DBPool.Exec(context.Background(), "select risky_public.delete_asset(fn_asset_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateAsset(assetInput types.Asset) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`select risky_public.update_asset(
			fn_asset_id => $1,
			fn_name => $2,
			fn_description => $3,
			fn_business_id => $4)`,
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
