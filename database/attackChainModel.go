package database

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type AttackChainModel struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
	ThreatID    uuid.UUID     `json:"assetId" db:"asset_id"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at"`
}

func (m *DBManager) GetAttackChains(businessID string) (attackChainOutput []AttackChainModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, capability_id, vulnerability_id, business_id, complexity, asset_id, created_at FROM risky_public.attackChains(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	attackChainOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[AttackChainModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetAttackChain(id string) (attackChainOutput AttackChainModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, capability_id, vulnerability_id, business_id, complexity, asset_id, created_at FROM risky_public.get_attack_chain(fn_attack_chain_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	attackChainOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[AttackChainModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteAttackChain(id string) (err error) {

	_, err = m.DBPool.Query(context.Background(), "select risky_public.delete_attack_chain(fn_attack_chain_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateAttackChain(attackChainInput AttackChainModel) (attackChainId string, err error) {

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

func (m *DBManager) UpdateAttackChain(attackChainInput AttackChainModel) (err error) {

	_, err = m.DBPool.Query(context.Background(),
		`select risky_public.update_attack_chain(
			fn_attack_chain_id => $1
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
