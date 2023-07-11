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

func TestGetImpacts(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	impacts, _ := dbManager.GetImpacts(businessId)

	for _, impact := range impacts {
		assert.Equal(t, impact.BusinessID.String(), businessId)
	}
}

func TestGetImpact(t *testing.T) {
	var impactId = "a507c9e8-8f46-4dee-aa69-5adece3e4372"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	impact, _ := dbManager.GetImpact(impactId)

	assert.Equal(t, impact.ID.String(), impactId)
}

func TestCreateImpact(t *testing.T) {
	var threatId = uuid.MustParse("f56d66b6-2543-435b-84da-51cdab340a01")
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	impactInput := types.Impact{Name: "test", BusinessID: uuid.MustParse(businessId), ThreatID: threatId}
	impactId, err := dbManager.CreateImpact(impactInput)

	assert.Equal(t, err, nil)

	impact, err := dbManager.GetImpact(impactId)

	assert.Equal(t, err, nil)

	assert.Equal(t, impact.ID.String(), impactId)
}

func TestDeleteImpact(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	impactInput := types.Impact{Name: "test", BusinessID: uuid.MustParse(businessId)}
	impactId, _ := dbManager.CreateImpact(impactInput)

	err = dbManager.DeleteImpact(impactId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetImpact(impactId)

	assert.NotEqual(t, err, nil)

}

func TestUpdateImpact(t *testing.T) {
	var threatId = uuid.MustParse("a8990e0a-1166-4097-9f1a-852d02d9ab19")

	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createImpactInput := types.Impact{Name: "test", BusinessID: uuid.MustParse(businessId), ThreatID: threatId}
	impactId, _ := dbManager.CreateImpact(createImpactInput)

	updateImpactInput := createImpactInput

	updateImpactInput.Name = "test2"
	updateImpactInput.ID = uuid.MustParse(impactId)

	err = dbManager.UpdateImpact(updateImpactInput)

	assert.Equal(t, err, nil)

	updatedImpact, _ := dbManager.GetImpact(impactId)

	assert.Equal(t, updateImpactInput.Name, updatedImpact.Name)
}
