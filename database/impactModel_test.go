package database_test

import (
	"context"
	"go-risky/database"
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
		assert.IsEqual(impact.BusinessID.String(), businessId)
	}
}

func TestGetImpact(t *testing.T) {
	var impactId = "535705bc-fddb-4e2a-8c1c-196755ce16b6"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	impact, _ := dbManager.GetImpact(impactId)

	assert.IsEqual(impact.ID.String(), impactId)
}

func TestDeleteImpact(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	impactInput := database.ImpactModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	impactId, _ := dbManager.CreateImpact(impactInput)

	err = dbManager.DeleteImpact(impactId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetImpact(impactId)

	assert.NotEqual(t, err, nil)

}

func TestCreateImpact(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	impactInput := database.ImpactModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	impactId, err := dbManager.CreateImpact(impactInput)

	assert.Equal(t, err, nil)

	impact, err := dbManager.GetImpact(impactId)

	assert.Equal(t, err, nil)

	assert.Equal(t, impact.ID.String(), impactId)
}

func TestUpdateImpact(t *testing.T) {

	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createImpactInput := database.ImpactModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	impactId, _ := dbManager.CreateImpact(createImpactInput)

	updateImpactInput := createImpactInput

	updateImpactInput.Name = "test2"
	updateImpactInput.ID = uuid.MustParse(impactId)

	err = dbManager.UpdateImpact(updateImpactInput)

	assert.Equal(t, err, nil)

	updatedImpact, _ := dbManager.GetImpact(impactId)

	assert.Equal(t, updateImpactInput.Name, updatedImpact.Name)
}
