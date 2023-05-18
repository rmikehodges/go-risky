package database

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type ThreatModel struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at"`
}

func (m *DBManager) GetThreats(businessID string) (threatOutput []ThreatModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id,created_at FROM risky_public.threats(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	threatOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[ThreatModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetThreat(id string) (threatOutput ThreatModel, err error) {

	rows, err := m.DBPool.Query(context.Background(), "select id,name, description,business_id, created_at FROM risky_public.get_threat(fn_threat_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	threatOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[ThreatModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteThreat(id string) (err error) {

	_, err = m.DBPool.Query(context.Background(), "select risky_public.delete_threat(fn_threat_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateThreat(threatInput ThreatModel) (err error) {

	_, err = m.DBPool.Query(context.Background(),
		`select risky_public.create_threat(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3)`,
		threatInput.Name,
		threatInput.Description,
		threatInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateThreat(threatInput ThreatModel) (err error) {

	_, err = m.DBPool.Query(context.Background(),
		`select risky_public.update_threat(
			fn_threat_id => $1
			fn_name => $2, 
			fn_description => $3
			fn_business_id => $4)`,
		threatInput.ID,
		threatInput.Name,
		threatInput.Description,
		threatInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
