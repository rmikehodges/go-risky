package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetAttackChains(businessID string) (attackChainOutput []types.AttackChain, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, threat_id, created_at FROM risky_public.attack_chains(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	attackChainOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.AttackChain])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetAttackChain(id string) (attackChainOutput types.AttackChain, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, threat_id, business_id, created_at FROM risky_public.get_attack_chain(fn_attack_chain_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	attackChainOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.AttackChain])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteAttackChain(id string) (err error) {

	_, err = m.DBPool.Exec(context.Background(), "select risky_public.delete_attack_chain(fn_attack_chain_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateAttackChain(attackChainInput types.AttackChain) (attackChainId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`select risky_public.create_attack_chain(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3,  
			fn_threat_id => $4)`,
		attackChainInput.Name,
		attackChainInput.Description,
		attackChainInput.BusinessID,
		attackChainInput.ThreatID).Scan(&attackChainId)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateAttackChain(attackChainInput types.AttackChain) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`select risky_public.update_attack_chain(
			fn_attack_chain_id => $1,
			fn_name => $2, 
			fn_description => $3, 
			fn_business_id => $4, 
			fn_threat_id => $5)`,
		attackChainInput.ID,
		attackChainInput.Name,
		attackChainInput.Description,
		attackChainInput.BusinessID,
		attackChainInput.ThreatID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
