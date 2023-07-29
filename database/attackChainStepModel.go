package database

import (
	"context"
	"fmt"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetAttackChainSteps(businessId string, attackChainId string, actionId string) (attackChainStepOutput []types.AttackChainStep, err error) {
	var rows pgx.Rows
	if attackChainId != "" {
		rows, err = m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.attack_chain_step WHERE business_id = $1 AND attack_chain_id = $2;", businessId, attackChainId)
		if err != nil {
			log.Println(err)
			return
		}
	} else if actionId != "" && attackChainId == "" {
		rows, err = m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.attack_chain_step WHERE business_id = $1 AND action_id = $2;", businessId, actionId)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		rows, err = m.DBPool.Query(context.Background(), " SELECT * FROM risky_public.attack_chain_step WHERE business_id = $1", businessId)
		if err != nil {
			log.Println(err)
			return
		}
	}

	attackChainStepOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.AttackChainStep])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetAttackChainStep(attackChainStepId string) (attackChainStepOutput types.AttackChainStep, err error) {

	rows, err := m.DBPool.Query(context.Background(), `SELECT * FROM risky_public.attack_chain_step WHERE id = $1;`, attackChainStepId)
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

	attackChainStep, err := m.GetAttackChainStep(attackChainStepId)
	if err != nil {
		log.Println(err)
		return
	}

	conn, err := m.DBPool.Acquire(context.Background())

	//Use transactions here
	tx, err := conn.Begin(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(),
		`UPDATE risky_public.attack_chain_step 
	SET next_step = $2 WHERE id = $1;`,
		attackChainStep.PreviousStep,
		attackChainStep.NextStep)

	if err != nil {
		log.Println(err)
		return
	}

	_, err = tx.Exec(context.Background(),
		`UPDATE risky_public.attack_chain_step 
	SET previous_step = $2 WHERE id = $1;`,
		attackChainStep.NextStep,
		attackChainStep.PreviousStep)

	_, err = tx.Exec(context.Background(),
		`DELETE FROM risky_public.attack_chain_step WHERE id = $1;`,
		attackChainStep.ID)

	err = tx.Commit(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateAttackChainStep(attackChainStepInput types.AttackChainStep) (attackChainStepId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`INSERT INTO 
		risky_public.attack_chain_step(attack_chain_id, action_id, asset_id, next_step, previous_step, business_id) 
		values($1, $2, $3,$4, $5 , $6) RETURNING id;
		`,
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
		`UPDATE risky_public.attack_chain_step SET 
		attack_chain_id = $2, 
		action_id = $3, 
		asset_id = $4, 
		next_step = $5, 
		previous_step = $6, 
		business_id = $7 
		WHERE id = $1;`,
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
