package database_test

import (
	"context"
	"go-risky/database"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var attackChainId = "535705bc-fddb-4e2a-8c1c-196755ce16b6"

func TestGetAttackChains(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	attackChains, _ := dbManager.GetAttackChains(businessId)

	for _, attackChain := range attackChains {
		assert.Equal(t, attackChain.BusinessID.String(), businessId)
	}
}

func TestGetAttackChain(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	attackChain, _ := dbManager.GetAttackChain(attackChainId)

	assert.Equal(t, attackChain.ID.String(), attackChainId)
}

func TestDeleteAttackChain(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	attackChainInput := database.AttackChainModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	attackChainId, _ := dbManager.CreateAttackChain(attackChainInput)

	err = dbManager.DeleteAttackChain(attackChainId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetAttackChain(attackChainId)

	assert.NotEqual(t, err, nil)

}

func TestCreateAttackChain(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	attackChainInput := database.AttackChainModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	attackChainId, err := dbManager.CreateAttackChain(attackChainInput)

	assert.Equal(t, err, nil)

	attackChain, err := dbManager.GetAttackChain(attackChainId)

	assert.Equal(t, err, nil)

	assert.Equal(t, attackChain.ID.String(), attackChainId)
}

func TestUpdateAttackChain(t *testing.T) {

	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createAttackChainInput := database.AttackChainModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	attackChainId, _ := dbManager.CreateAttackChain(createAttackChainInput)

	updateAttackChainInput := createAttackChainInput
	updateAttackChainInput.Name = "test2"
	updateAttackChainInput.ID = uuid.MustParse(attackChainId)

	err = dbManager.UpdateAttackChain(updateAttackChainInput)

	assert.Equal(t, err, nil)

	updatedAttackChain, _ := dbManager.GetAttackChain(attackChainId)

	assert.Equal(t, updateAttackChainInput.Name, updatedAttackChain.Name)
}
