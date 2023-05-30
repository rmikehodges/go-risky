package database_test

import (
	"context"
	"go-risky/database"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestGetActions(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	actions, _ := dbManager.GetActions(businessId)

	for _, action := range actions {
		assert.Equal(t, action.BusinessID.String(), businessId)
	}
}

func TestGetAction(t *testing.T) {
	var actionId = "535705bc-fddb-4e2a-8c1c-196755ce16b6"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	action, _ := dbManager.GetAction(actionId)

	assert.Equal(t, action.ID.String(), actionId)
}

func TestCreateAction(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	actionInput := database.ActionModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	actionId, err := dbManager.CreateAction(actionInput)

	assert.Equal(t, err, nil)

	action, err := dbManager.GetAction(actionId)

	assert.Equal(t, err, nil)

	assert.Equal(t, action.ID.String(), actionId)
}

func TestDeleteAction(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	actionInput := database.ActionModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	actionId, _ := dbManager.CreateAction(actionInput)

	err = dbManager.DeleteAction(actionId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetAction(actionId)

	assert.NotEqual(t, err, nil)

}

func TestUpdateAction(t *testing.T) {

	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createActionInput := database.ActionModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	actionId, _ := dbManager.CreateAction(createActionInput)

	createActionInput.Name = "test2"
	createActionInput.ID = uuid.MustParse(actionId)

	updateActionInput := createActionInput

	err = dbManager.UpdateAction(updateActionInput)

	assert.Equal(t, err, nil)

	updatedAction, _ := dbManager.GetAction(actionId)

	assert.Equal(t, updateActionInput.Name, updatedAction.Name)
}
