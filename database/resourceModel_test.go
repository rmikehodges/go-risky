package database_test

import (
	"context"
	"go-risky/database"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestGetResources(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	resources, _ := dbManager.GetResources(businessId)

	for _, resource := range resources {
		assert.Equal(t, resource.BusinessID.String(), businessId)
	}
}

func TestGetResource(t *testing.T) {
	var resourceId = "82021d2b-a7df-4a22-95b6-19c1039db441"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	resource, _ := dbManager.GetResource(resourceId)

	assert.Equal(t, resource.ID.String(), resourceId)
}

func TestCreateResource(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	resourceInput := database.ResourceModel{Name: "test", Description: "test", Cost: 100, Unit: "dollar", ResourceType: "CASH", Total: 100000, BusinessID: uuid.MustParse(businessId)}
	resourceId, err := dbManager.CreateResource(resourceInput)

	assert.Equal(t, err, nil)

	resource, err := dbManager.GetResource(resourceId)

	assert.Equal(t, err, nil)

	assert.Equal(t, resource.ID.String(), resourceId)
}

func TestDeleteResource(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	resourceInput := database.ResourceModel{Name: "test", Description: "test", Cost: 100, Unit: "dollar", ResourceType: "CASH", Total: 100000, BusinessID: uuid.MustParse(businessId)}
	resourceId, _ := dbManager.CreateResource(resourceInput)

	err = dbManager.DeleteResource(resourceId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetResource(resourceId)

	assert.NotEqual(t, err, nil)

}

func TestUpdateResource(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createResourceInput := database.ResourceModel{Name: "test", Description: "test", Cost: 100, Unit: "dollar", ResourceType: "CASH", Total: 100000, BusinessID: uuid.MustParse(businessId)}
	resourceId, _ := dbManager.CreateResource(createResourceInput)

	updateResourceInput := createResourceInput
	updateResourceInput.Name = "test2"
	updateResourceInput.ID = uuid.MustParse(resourceId)

	err = dbManager.UpdateResource(updateResourceInput)

	assert.Equal(t, err, nil)

	updatedResource, _ := dbManager.GetResource(resourceId)

	assert.Equal(t, updateResourceInput.ID, updatedResource.ID)
}
