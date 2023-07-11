package database_test

import (
	"context"
	"go-risky/database"
	"go-risky/types"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestGetThreats(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	threats, _ := dbManager.GetThreats(businessId)

	for _, threat := range threats {
		assert.Equal(t, threat.BusinessID.String(), businessId)
	}
}

func TestGetThreat(t *testing.T) {
	var threatId = "f56d66b6-2543-435b-84da-51cdab340a01"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	threat, _ := dbManager.GetThreat(threatId)

	assert.Equal(t, threat.ID.String(), threatId)
}

func TestCreateThreat(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	threatInput := types.Threat{Name: "test", BusinessID: uuid.MustParse(businessId)}
	threatId, err := dbManager.CreateThreat(threatInput)

	assert.Equal(t, err, nil)

	threat, err := dbManager.GetThreat(threatId)

	assert.Equal(t, err, nil)

	assert.Equal(t, threat.ID.String(), threatId)
}

func TestDeleteThreat(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	threatInput := types.Threat{Name: "test", BusinessID: uuid.MustParse(businessId)}
	threatId, _ := dbManager.CreateThreat(threatInput)

	err = dbManager.DeleteThreat(threatId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetThreat(threatId)

	assert.NotEqual(t, err, nil)

}

func TestUpdateThreat(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createThreatInput := types.Threat{Name: "test", BusinessID: uuid.MustParse(businessId)}
	threatId, _ := dbManager.CreateThreat(createThreatInput)

	updateThreatInput := createThreatInput
	updateThreatInput.Name = "test2"
	updateThreatInput.ID = uuid.MustParse(threatId)

	err = dbManager.UpdateThreat(updateThreatInput)

	assert.Equal(t, err, nil)

	updatedThreat, _ := dbManager.GetThreat(threatId)

	assert.Equal(t, updateThreatInput.ID, updatedThreat.ID)
}
