package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetAttackChains(businessID string) (attackChainOutput []types.AttackChain, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.attack_chain WHERE business_id = $1", businessID)
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

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.attack_chain WHERE id = $1", id)
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

	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_public.attack_chain WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateAttackChain(attackChainInput types.AttackChain) (attackChainId string, err error) {

	err = m.DBPool.QueryRow(context.Background(),
		`INSERT INTO risky_public.attack_chain(name, description, business_id, threat_id) VALUES ($1, $2, $3, $4) RETURNING id;`,
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
		`UPDATE risky_public.attack_chain SET name = $2, description = $3, business_id = $4, threat_id = $5 WHERE id = $1;`,
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
