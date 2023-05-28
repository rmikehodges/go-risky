package database_test

import (
	"context"
	"go-risky/database"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestGetLiabilities(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	liabilities, _ := dbManager.GetLiabilities(businessId)

	for _, liability := range liabilities {
		assert.IsEqual(liability.BusinessID.String(), businessId)
	}
}

func TestGetLiability(t *testing.T) {
	var liabilityId = "535705bc-fddb-4e2a-8c1c-196755ce16b6"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	liability, _ := dbManager.GetLiability(liabilityId)

	assert.IsEqual(liability.ID.String(), liabilityId)
}

func TestDeleteLiability(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	liabilityInput := database.LiabilityModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	liabilityId, _ := dbManager.CreateLiability(liabilityInput)

	err = dbManager.DeleteLiability(liabilityId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetLiability(liabilityId)

	assert.NotEqual(t, err, nil)

}

func TestCreateLiability(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	liabilityInput := database.LiabilityModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	liabilityId, err := dbManager.CreateLiability(liabilityInput)

	assert.Equal(t, err, nil)

	liability, err := dbManager.GetLiability(liabilityId)

	assert.Equal(t, err, nil)

	assert.Equal(t, liability.ID.String(), liabilityId)
}

func TestUpdateLiability(t *testing.T) {

	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createLiabilityInput := database.LiabilityModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	liabilityId, _ := dbManager.CreateLiability(createLiabilityInput)

	updateLiabilityInput := createLiabilityInput
	updateLiabilityInput.Name = "test2"
	updateLiabilityInput.ID = uuid.MustParse(liabilityId)

	err = dbManager.UpdateLiability(updateLiabilityInput)

	assert.Equal(t, err, nil)

	updatedLiability, _ := dbManager.GetLiability(liabilityId)

	assert.Equal(t, updateLiabilityInput.Name, updatedLiability.Name)
}
