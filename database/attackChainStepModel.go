package database

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// Create type AttackChainStepModel based on table defintition risky_public.attack_chain_action in database/tables.sql
type AttackChainStepModel struct {
	BusinessID    uuid.UUID `db:"business_id"`
	ActionID      uuid.UUID `db:"action_id"`
	AttackChainID uuid.UUID `db:"attack_chain_id"`
	Position      int       `db:"position"`
	CreatedAt     time.Time `db:"created_at"`
}

func (m *DBManager) GetAttackChainSteps(attackChainId string, businessId string) (actionOutput []AttackChainStepModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select attack_chain_id, action_id,position,business_id, created_at FROM risky_public.attack_chain_steps(fn_business_id => $1, fn_attack_chain_id => $2)", businessId, attackChainId)
	if err != nil {
		log.Println(err)
		return
	}

	actionOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[AttackChainStepModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetAttackChainStep(actionId string, attackChainId string) (actionOutput AttackChainStepModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select  attack_chain_id, action_id,position,business_id, FROM risky_public.get_attack_chain_step(fn_attack_chain_id => $1, fn_action_id => $2)", attackChainId, actionId)
	if err != nil {
		log.Println(err)
		return
	}

	actionOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[AttackChainStepModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteAttackChainStep(actionId string, attackChainId string) (err error) {

	_, err = m.DBPool.Query(context.Background(),
		"select risky_public.delete_attack_chain_step(fn_attack_chain_id => $1, fn_action_id => $2)", attackChainId, actionId)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateAttackChainStep(attackChainStepInput AttackChainStepModel) (actionId string, attackChainId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`select risky_public.create_attack_chain_step(
			fn_attack_chain_id => $1, 
			fn_action_id => $2, 
			fn_position => $3
			fn_business_id => $4)`,
		attackChainStepInput.AttackChainID,
		attackChainStepInput.ActionID,
		attackChainStepInput.Position,
		attackChainStepInput.BusinessID).Scan(&attackChainId, &actionId)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateAttackChainStep(attackChainStepInput AttackChainStepModel) (err error) {

	_, err = m.DBPool.Query(context.Background(),
		`select risky_public.update_attack_chain_step(
			fn_attack_chain_id => $1, 
			fn_action_id => $2, 
			fn_position => $3
			fn_business_id => $4)`,
		attackChainStepInput.AttackChainID,
		attackChainStepInput.ActionID,
		attackChainStepInput.Position,
		attackChainStepInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
