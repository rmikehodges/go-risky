package database_test

import (
	"context"
	"fmt"
	"go-risky/database"
	"go-risky/types"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestGetAttackChainSteps(t *testing.T) {
	var attackChainId = "20036fa3-45c6-47b2-a343-f88bcd4f5e07"
	var actionId = ""
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	attackChainSteps, _ := dbManager.GetAttackChainSteps(attackChainId, businessId, actionId)

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
	attackChainStepInput := types.AttackChainStep{ActionID: uuid.MustParse(actionId), AttackChainID: uuid.MustParse(attackChainId), AssetID: &assetId, BusinessID: uuid.MustParse(businessId)}
	createdAttackChainStepId, _ := dbManager.CreateAttackChainStep(attackChainStepInput)
	attackChainStep, _ := dbManager.GetAttackChainStep(createdAttackChainStepId)

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
	attackChainStepInput := types.AttackChainStep{ActionID: actionId, AttackChainID: attackChainId, AssetID: &assetId, BusinessID: uuid.MustParse(businessId)}
	attackChainStepId, _ := dbManager.CreateAttackChainStep(attackChainStepInput)

	assert.Equal(t, err, nil)

	attackChainStep, err := dbManager.GetAttackChainStep(attackChainStepId)

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
	attackChainStepInput := types.AttackChainStep{ActionID: actionId, AttackChainID: attackChainId, AssetID: &assetId, BusinessID: uuid.MustParse(businessId)}
	attackChainStepId, err := dbManager.CreateAttackChainStep(attackChainStepInput)
	fmt.Println(attackChainStepId)
	if err != nil {
		panic(err)
	}

	err = dbManager.DeleteAttackChainStep(attackChainStepId)

	assert.Equal(t, err, nil)

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
	attackChainStepInput := types.AttackChainStep{ActionID: actionId, AttackChainID: attackChainId, AssetID: &assetId, BusinessID: uuid.MustParse(businessId)}
	attackChainStepId, _ := dbManager.CreateAttackChainStep(attackChainStepInput)

	attackChainStepInput.ID = uuid.MustParse(attackChainStepId)
	attackChainStepInput.ActionID = uuid.MustParse("73088b69-dbc2-4f93-bf4a-de292af69102")

	err = dbManager.UpdateAttackChainStep(attackChainStepInput)

	assert.Equal(t, err, nil)

	updatedAttackChainStep, err := dbManager.GetAttackChainStep(attackChainStepId)

	assert.Equal(t, err, nil)

	assert.Equal(t, attackChainStepInput.ActionID, updatedAttackChainStep.ActionID)
	assert.Equal(t, attackChainStepInput.AttackChainID, updatedAttackChainStep.AttackChainID)
	// assert.Equal(t, attackChainStepInput.Position, updatedAttackChainStep.Position)
}
