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
	var attackChainId = "20036fa3-45c6-47b2-a343-f88bcd4f5e07"
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

func TestCreateAttackChain(t *testing.T) {
	var threatId = "f56d66b6-2543-435b-84da-51cdab340a01"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	attackChainInput := types.AttackChain{Name: "test", ThreatID: uuid.MustParse(threatId), BusinessID: uuid.MustParse(businessId)}
	attackChainId, err := dbManager.CreateAttackChain(attackChainInput)

	assert.Equal(t, err, nil)

	attackChain, err := dbManager.GetAttackChain(attackChainId)

	assert.Equal(t, err, nil)

	assert.Equal(t, attackChain.ID.String(), attackChainId)
}

func TestDeleteAttackChain(t *testing.T) {
	var threatId = "f56d66b6-2543-435b-84da-51cdab340a01"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	attackChainInput := types.AttackChain{Name: "test", ThreatID: uuid.MustParse(threatId), BusinessID: uuid.MustParse(businessId)}
	attackChainId, _ := dbManager.CreateAttackChain(attackChainInput)

	err = dbManager.DeleteAttackChain(attackChainId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetAttackChain(attackChainId)

	assert.NotEqual(t, err, nil)

}

func TestUpdateAttackChain(t *testing.T) {
	var threatId = "f56d66b6-2543-435b-84da-51cdab340a01"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createAttackChainInput := types.AttackChain{Name: "test", ThreatID: uuid.MustParse(threatId), BusinessID: uuid.MustParse(businessId)}

	attackChainId, _ := dbManager.CreateAttackChain(createAttackChainInput)

	updateAttackChainInput := types.AttackChain{ID: uuid.MustParse(attackChainId), Name: "test2", ThreatID: uuid.MustParse(threatId), BusinessID: uuid.MustParse(businessId)}

	err = dbManager.UpdateAttackChain(updateAttackChainInput)

	assert.Equal(t, err, nil)

	updatedAttackChain, err := dbManager.GetAttackChain(updateAttackChainInput.ID.String())

	assert.Equal(t, err, nil)
	assert.Equal(t, updateAttackChainInput.Name, updatedAttackChain.Name)
}
