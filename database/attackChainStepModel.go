package database

import (
	"context"
	"fmt"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetAttackChainSteps(businessId string, attackChainId string) (attackChainStepOutput []types.AttackChainStep, err error) {
	rows, err := m.DBPool.Query(context.Background(), "select id, attack_chain_id, action_id, detection_id, mitigation_id, asset_id, next_step, previous_step,business_id, created_at FROM risky_public.attack_chain_steps(fn_business_id => $1, fn_attack_chain_id => $2)", businessId, attackChainId)
	if err != nil {
		log.Println(err)
		return
	}

	attackChainStepOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.AttackChainStep])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetAttackChainStep(attackChainStepId string) (attackChainStepOutput types.AttackChainStep, err error) {

	rows, err := m.DBPool.Query(context.Background(), `select  id, 
	attack_chain_id, action_id, asset_id, next_step, previous_step ,business_id, detection_id, mitigation_id,created_at 
	FROM 
	risky_public.get_attack_chain_step(fn_attack_chain_step_id => $1)`,
		attackChainStepId)
	if err != nil {
		log.Println(err)
		return
	}

	attackChainStepOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.AttackChainStep])
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

func (m *DBManager) CreateAttackChainStep(attackChainStepInput types.AttackChainStep) (attackChainStepId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`select * FROM risky_public.create_attack_chain_step(
			fn_attack_chain_id => $1, 
			fn_action_id => $2,
			fn_asset_id => $3,
			fn_next_step => $4,
			fn_previous_step => $5,
			fn_business_id => $6)`,
		attackChainStepInput.AttackChainID,
		attackChainStepInput.ActionID,
		attackChainStepInput.AssetID,
		attackChainStepInput.NextStep,
		attackChainStepInput.PreviousStep,
		attackChainStepInput.BusinessID).Scan(&attackChainStepId)
	if err != nil {
		log.Printf("CreateAttackChainStep Error: %s", err)
		return
	}

	return
}

func (m *DBManager) UpdateAttackChainStep(attackChainStepInput types.AttackChainStep) (err error) {
	fmt.Println(attackChainStepInput.ID)
	_, err = m.DBPool.Exec(context.Background(),
		`select risky_public.update_attack_chain_step(
			fn_attack_chain_step_id => $1, 
			fn_attack_chain_id => $2, 
			fn_action_id => $3, 
			fn_asset_id => $4,
			fn_next_step => $5,
			fn_previous_step => $6,
			fn_business_id => $7)`,
		attackChainStepInput.ID,
		attackChainStepInput.AttackChainID,
		attackChainStepInput.ActionID,
		attackChainStepInput.AssetID,
		attackChainStepInput.NextStep,
		attackChainStepInput.PreviousStep,
		attackChainStepInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
