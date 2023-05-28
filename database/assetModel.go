package database

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type AssetModel struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at"`
}

func (m *DBManager) GetAssets(businessID string) (assetOutput []AssetModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, created_at FROM risky_public.assets(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	assetOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[AssetModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetAsset(id string) (assetOutput AssetModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, created_at FROM risky_public.get_asset(fn_asset_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	assetOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[AssetModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateAsset(assetInput AssetModel) (assetId string, err error) {

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

func (m *DBManager) UpdateAsset(assetInput AssetModel) (err error) {

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
