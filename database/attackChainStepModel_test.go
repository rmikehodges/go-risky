package database_test

import (
	"context"
	"go-risky/database"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestGetAttackChainSteps(t *testing.T) {
	var attackChainId = "20036fa3-45c6-47b2-a343-f88bcd4f5e07"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	attackChainSteps, _ := dbManager.GetAttackChainSteps(attackChainId, businessId)

	for _, attackChainStep := range attackChainSteps {
		assert.Equal(t, attackChainStep.BusinessID.String(), businessId)
	}
}

func TestGetAttackChainStep(t *testing.T) {
	var actionId = "cdf5e362-da33-48aa-8d93-a4358b05789e"
	var attackChainId = "20036fa3-45c6-47b2-a343-f88bcd4f5e07"
	var assetId = uuid.MustParse("1c1c31ce-70aa-47aa-a0e3-fdeabcb4957c")
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	attackChainStepInput := database.AttackChainStepModel{ActionID: uuid.MustParse(actionId), AttackChainID: uuid.MustParse(attackChainId), AssetID: &assetId, BusinessID: uuid.MustParse(businessId), Position: 1}
	createdAttackChainStep, _ := dbManager.CreateAttackChainStep(attackChainStepInput)
	attackChainStep, _ := dbManager.GetAttackChainStep(createdAttackChainStep.ID.String())

	assert.Equal(t, attackChainStep.ActionID.String(), actionId)
	assert.Equal(t, attackChainStep.AttackChainID.String(), attackChainId)

}

func TestCreateAttackChainStep(t *testing.T) {
	var actionId = uuid.MustParse("cdf5e362-da33-48aa-8d93-a4358b05789e")
	var attackChainId = uuid.MustParse("20036fa3-45c6-47b2-a343-f88bcd4f5e07")
	var assetId = uuid.MustParse("465804b9-e5aa-49e1-b844-61ba3d928b84")

	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	attackChainStepInput := database.AttackChainStepModel{ActionID: actionId, AttackChainID: attackChainId, AssetID: &assetId, BusinessID: uuid.MustParse(businessId), Position: 1}
	attackChainStep, _ := dbManager.CreateAttackChainStep(attackChainStepInput)

	assert.Equal(t, err, nil)

	attackChainStep, err = dbManager.GetAttackChainStep(attackChainStep.ID.String())

	assert.Equal(t, err, nil)

	assert.Equal(t, attackChainStepInput.ActionID, attackChainStep.ActionID)
	assert.Equal(t, attackChainStepInput.AttackChainID, attackChainStep.AttackChainID)
}

func TestDeleteAttackChainStep(t *testing.T) {
	var actionId = uuid.MustParse("cdf5e362-da33-48aa-8d93-a4358b05789e")
	var attackChainId = uuid.MustParse("20036fa3-45c6-47b2-a343-f88bcd4f5e07")
	var assetId = uuid.MustParse("465804b9-e5aa-49e1-b844-61ba3d928b84")
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	attackChainStepInput := database.AttackChainStepModel{ActionID: actionId, AttackChainID: attackChainId, AssetID: &assetId, BusinessID: uuid.MustParse(businessId), Position: 1}
	attackChainStep, _ := dbManager.CreateAttackChainStep(attackChainStepInput)

	err = dbManager.DeleteAttackChainStep(attackChainStep.ID.String())

	assert.Equal(t, err, nil)

	_, err = dbManager.GetAttackChainStep(attackChainStep.ID.String())

	assert.NotEqual(t, err, nil)

}

func TestUpdateAttackChainStep(t *testing.T) {
	var actionId = uuid.MustParse("cdf5e362-da33-48aa-8d93-a4358b05789e")
	var attackChainId = uuid.MustParse("20036fa3-45c6-47b2-a343-f88bcd4f5e07")
	var assetId = uuid.MustParse("465804b9-e5aa-49e1-b844-61ba3d928b84")
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	attackChainStepInput := database.AttackChainStepModel{ActionID: actionId, AttackChainID: attackChainId, AssetID: &assetId, BusinessID: uuid.MustParse(businessId), Position: 1}
	attackChainStep, _ := dbManager.CreateAttackChainStep(attackChainStepInput)

	updateAttackChainStepInput := attackChainStep
	updateAttackChainStepInput.AttackChainID = (attackChainStep.AttackChainID)
	updateAttackChainStepInput.ActionID = uuid.MustParse("73088b69-dbc2-4f93-bf4a-de292af69102")
	updateAttackChainStepInput.Position = 1

	err = dbManager.UpdateAttackChainStep(updateAttackChainStepInput)

	assert.Equal(t, err, nil)

	updatedAttackChainStep, err := dbManager.GetAttackChainStep(attackChainStep.ID.String())

	assert.Equal(t, err, nil)

	assert.Equal(t, updateAttackChainStepInput.ActionID, updatedAttackChainStep.ActionID)
	assert.Equal(t, updateAttackChainStepInput.AttackChainID, updatedAttackChainStep.AttackChainID)
	assert.Equal(t, updateAttackChainStepInput.Position, updatedAttackChainStep.Position)
}
