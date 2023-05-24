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
var actionId = "535705bc-fddb-4e2a-8c1c-196755ce16b6"

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
		assert.IsEqual(action.BusinessID.String(), businessId)
	}
}

func TestGetAction(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	action, _ := dbManager.GetAction(actionId)

	assert.IsEqual(action.ID.String(), actionId)
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
	createActionOutput, _ := dbManager.CreateAction(actionInput)

	tempActionId := createActionOutput.ID.String()

	err = dbManager.DeleteAction(tempActionId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetAction(tempActionId)

	assert.NotEqual(t, err, nil)

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
	createActionOutput, err := dbManager.CreateAction(actionInput)

	assert.Equal(t, err, nil)

	action, err := dbManager.GetAction(createActionOutput.ID.String())

	assert.Equal(t, err, nil)

	assert.Equal(t, action, createActionOutput)
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
	createActionOutput, _ := dbManager.CreateAction(createActionInput)

	createActionOutput.Name = "test2"

	updateActionInput := createActionOutput

	err = dbManager.UpdateAction(updateActionInput)

	assert.Equal(t, err, nil)

	updatedAction, _ := dbManager.GetAction(createActionOutput.ID.String())

	assert.Equal(t, updateActionInput.Name, updatedAction.Name)
}
