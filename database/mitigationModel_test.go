package database_test

import (
	"context"
	"go-risky/database"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestGetMitigations(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	mitigations, _ := dbManager.GetMitigations(businessId)

	for _, mitigation := range mitigations {
		assert.IsEqual(mitigation.BusinessID.String(), businessId)
	}
}

func TestGetMitigation(t *testing.T) {
	var mitigationId = "535705bc-fddb-4e2a-8c1c-196755ce16b6"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	mitigation, _ := dbManager.GetMitigation(mitigationId)

	assert.IsEqual(mitigation.ID.String(), mitigationId)
}

func TestDeleteMitigation(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	mitigationInput := database.MitigationModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	mitigationId, _ := dbManager.CreateMitigation(mitigationInput)

	err = dbManager.DeleteMitigation(mitigationId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetMitigation(mitigationId)

	assert.NotEqual(t, err, nil)

}

func TestCreateMitigation(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	mitigationInput := database.MitigationModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	mitigationId, err := dbManager.CreateMitigation(mitigationInput)

	assert.Equal(t, err, nil)

	mitigation, err := dbManager.GetMitigation(mitigationId)

	assert.Equal(t, err, nil)

	assert.Equal(t, mitigation.ID.String(), mitigationId)
}

func TestUpdateMitigation(t *testing.T) {

	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createMitigationInput := database.MitigationModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	mitigationId, _ := dbManager.CreateMitigation(createMitigationInput)

	updateMitigationInput := createMitigationInput
	updateMitigationInput.Name = "test2"
	updateMitigationInput.ID = uuid.MustParse(mitigationId)

	err = dbManager.UpdateMitigation(updateMitigationInput)

	assert.Equal(t, err, nil)

	updatedMitigation, _ := dbManager.GetMitigation(mitigationId)

	assert.Equal(t, updateMitigationInput.Name, updatedMitigation.Name)
}
