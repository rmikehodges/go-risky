package database_test

import (
	"context"
	"go-risky/database"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var assetId = "465804b9-e5aa-49e1-b844-61ba3d928b84"

func TestGetAssets(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	assets, _ := dbManager.GetAssets(businessId)

	for _, asset := range assets {
		assert.IsEqual(asset.BusinessID.String(), businessId)
	}
}

func TestGetAsset(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	asset, _ := dbManager.GetAsset(assetId)

	assert.IsEqual(asset.ID.String(), assetId)
}

func TestCreateAsset(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	assetInput := database.AssetModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	assetId, err := dbManager.CreateAsset(assetInput)

	assert.Equal(t, err, nil)

	asset, err := dbManager.GetAsset(assetId)

	assert.Equal(t, err, nil)

	assert.Equal(t, asset.ID.String(), assetId)
}

func TestDeleteAsset(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	assetInput := database.AssetModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	assetId, _ := dbManager.CreateAsset(assetInput)

	err = dbManager.DeleteAsset(assetId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetAsset(assetId)

	assert.NotEqual(t, err, nil)

}

func TestUpdateAsset(t *testing.T) {

	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createAssetInput := database.AssetModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	assetId, _ := dbManager.CreateAsset(createAssetInput)

	updateAssetInput := createAssetInput

	updateAssetInput.Name = "test2"
	updateAssetInput.ID = uuid.MustParse(assetId)

	err = dbManager.UpdateAsset(updateAssetInput)

	assert.Equal(t, err, nil)

	updatedAsset, _ := dbManager.GetAsset(assetId)

	assert.Equal(t, updateAssetInput.Name, updatedAsset.Name)
}
