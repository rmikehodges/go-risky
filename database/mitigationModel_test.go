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
		assert.Equal(t, mitigation.BusinessID.String(), businessId)
	}
}

func TestGetMitigation(t *testing.T) {
	var mitigationId = "ab6b6ddb-8a6e-4102-a900-1acca26a404b"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	mitigation, _ := dbManager.GetMitigation(mitigationId)

	assert.Equal(t, mitigation.ID.String(), mitigationId)
}

func TestCreateMitigation(t *testing.T) {
	var actionId = uuid.MustParse("535705bc-fddb-4e2a-8c1c-196755ce16b6")
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	mitigationInput := database.MitigationModel{Name: "test", BusinessID: uuid.MustParse(businessId), ActionID: &actionId}
	mitigationId, err := dbManager.CreateMitigation(mitigationInput)

	assert.Equal(t, err, nil)

	mitigation, err := dbManager.GetMitigation(mitigationId)

	assert.Equal(t, err, nil)

	assert.Equal(t, mitigation.ID.String(), mitigationId)
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

func TestUpdateMitigation(t *testing.T) {
	var actionId = uuid.MustParse("535705bc-fddb-4e2a-8c1c-196755ce16b6")
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createMitigationInput := database.MitigationModel{Name: "test", BusinessID: uuid.MustParse(businessId), ActionID: &actionId}
	mitigationId, _ := dbManager.CreateMitigation(createMitigationInput)

	updateMitigationInput := createMitigationInput
	updateMitigationInput.Name = "test2"
	updateMitigationInput.ID = uuid.MustParse(mitigationId)

	err = dbManager.UpdateMitigation(updateMitigationInput)

	assert.Equal(t, err, nil)

	updatedMitigation, _ := dbManager.GetMitigation(mitigationId)

	assert.Equal(t, updateMitigationInput.ID, updatedMitigation.ID)
}
