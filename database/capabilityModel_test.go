package database_test

import (
	"context"
	"go-risky/database"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestGetCapabilities(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	capabilities, _ := dbManager.GetCapabilities(businessId)

	for _, capability := range capabilities {
		assert.Equal(t, capability.BusinessID.String(), businessId)
	}
}

func TestGetCapability(t *testing.T) {
	var capabilityId = "8f506766-41c4-447b-a0ee-f43fc7fd1487"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	capability, _ := dbManager.GetCapability(capabilityId)

	assert.Equal(t, capability.ID.String(), capabilityId)
}

func TestCreateCapability(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	capabilityInput := database.CapabilityModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	capabilityId, err := dbManager.CreateCapability(capabilityInput)

	assert.Equal(t, err, nil)

	capability, err := dbManager.GetCapability(capabilityId)

	assert.Equal(t, err, nil)

	assert.Equal(t, capability.ID.String(), capabilityId)
}

func TestDeleteCapability(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	capabilityInput := database.CapabilityModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	capabilityId, _ := dbManager.CreateCapability(capabilityInput)

	err = dbManager.DeleteCapability(capabilityId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetCapability(capabilityId)

	assert.NotEqual(t, err, nil)

}

func TestUpdateCapability(t *testing.T) {

	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createCapabilityInput := database.CapabilityModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	capabilityId, _ := dbManager.CreateCapability(createCapabilityInput)

	updateCapabilityInput := createCapabilityInput
	updateCapabilityInput.Name = "test2"
	updateCapabilityInput.ID = uuid.MustParse(capabilityId)

	err = dbManager.UpdateCapability(updateCapabilityInput)

	assert.Equal(t, err, nil)

	updatedCapability, _ := dbManager.GetCapability(capabilityId)

	assert.Equal(t, updateCapabilityInput.Name, updatedCapability.Name)
}
