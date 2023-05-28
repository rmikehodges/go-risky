package database_test

import (
	"context"
	"go-risky/database"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var businessId = "23628819-59dd-45f3-8395-aceeca86bc9c"

func TestGetBusinesses(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	_, err = dbManager.GetBusinesses()

	assert.Equal(t, err, nil)

	// for _, business := range businesses {
	// 	assert.IsEqual(business.ID.String(), businessId)
	// }
}

func TestGetBusiness(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	business, _ := dbManager.GetBusiness(businessId)

	assert.IsEqual(business.ID.String(), businessId)
}

func TestDeleteBusiness(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	businessInput := database.BusinessModel{Name: "test", Revenue: 10000}
	businessId, _ := dbManager.CreateBusiness(businessInput)

	err = dbManager.DeleteBusiness(businessId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetBusiness(businessId)

	assert.NotEqual(t, err, nil)

}

func TestCreateBusiness(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	businessInput := database.BusinessModel{Name: "test", Revenue: 10000}
	businessId, err := dbManager.CreateBusiness(businessInput)

	assert.Equal(t, err, nil)

	business, err := dbManager.GetBusiness(businessId)

	assert.Equal(t, err, nil)

	assert.Equal(t, business.ID.String(), businessId)
}

func TestUpdateBusiness(t *testing.T) {

	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createBusinessInput := database.BusinessModel{Name: "test", Revenue: 10000}
	businessId, _ := dbManager.CreateBusiness(createBusinessInput)

	updateBusinessInput := createBusinessInput

	updateBusinessInput.Name = "test2"
	updateBusinessInput.ID = uuid.MustParse(businessId)

	err = dbManager.UpdateBusiness(updateBusinessInput)

	assert.Equal(t, err, nil)

	updatedBusiness, _ := dbManager.GetBusiness(businessId)

	assert.Equal(t, updateBusinessInput.Name, updatedBusiness.Name)
}
