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
	ID            uuid.UUID  `db:"id"`
	BusinessID    uuid.UUID  `db:"business_id"`
	ActionID      uuid.UUID  `db:"action_id"`
	AssetID       *uuid.UUID `db:"asset_id"`
	AttackChainID uuid.UUID  `db:"attack_chain_id"`
	Position      int        `db:"position"`
	DetectionID   uuid.UUID  `db:"detection_id"`
	MitigationID  *uuid.UUID `db:"mitigation_id"`
	CreatedAt     time.Time  `db:"created_at"`
}

func (m *DBManager) GetAttackChainSteps(attackChainId string, businessId string) (attackChainStepOutput []AttackChainStepModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select attack_chain_id, action_id, asset_id, position,business_id, created_at FROM risky_public.attack_chain_steps(fn_business_id => $1, fn_attack_chain_id => $, fn_attack_chain_id => $2)", businessId, attackChainId)
	if err != nil {
		log.Println(err)
		return
	}

	attackChainStepOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[AttackChainStepModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetAttackChainStep(attackChainStepId string) (attackChainStepOutput AttackChainStepModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), `select  id, 
	attack_chain_id, action_id, asset_id, position,business_id, detection_id, mitigation_id,created_at 
	FROM 
	risky_public.get_attack_chain_step(fn_attack_chain_step_id => $1)`,
		attackChainStepId)
	if err != nil {
		log.Println(err)
		return
	}

	attackChainStepOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[AttackChainStepModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteAttackChainStep(attackChainStepId string) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`select risky_public.delete_attack_chain_step(fn_attack_chain_step_id => $1)`, attackChainStepId)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateAttackChainStep(attackChainStepInput AttackChainStepModel) (attackChainStepOutput AttackChainStepModel, err error) {

	rows, err := m.DBPool.Query(context.Background(),
		`select id, 
		attack_chain_id, action_id, asset_id, position,business_id, detection_id, mitigation_id, created_at
		FROM risky_public.create_attack_chain_step(
			fn_attack_chain_id => $1, 
			fn_action_id => $2,
			fn_asset_id => $3 ,
			fn_position => $4,
			fn_business_id => $5)`,
		attackChainStepInput.AttackChainID,
		attackChainStepInput.ActionID,
		attackChainStepInput.AssetID,
		attackChainStepInput.Position,
		attackChainStepInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	attackChainStepOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[AttackChainStepModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateAttackChainStep(attackChainStepInput AttackChainStepModel) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`select risky_public.update_attack_chain_step(
			fn_attack_chain_step_id => $1, 
			fn_attack_chain_id => $2, 
			fn_action_id => $3, 
			fn_asset_id => $4,
			fn_position => $5,
			fn_business_id => $6)`,
		attackChainStepInput.ID,
		attackChainStepInput.AttackChainID,
		attackChainStepInput.ActionID,
		attackChainStepInput.AssetID,
		attackChainStepInput.Position,
		attackChainStepInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
