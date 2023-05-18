package database

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type DetectionModel struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
	ActionID    uuid.UUID     `json:"actionId" db:"action_id"`
	Implemented bool          `json:"complexity"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at"`
}

func (m *DBManager) GetDetections(businessID string) (detectionOutput []DetectionModel, err error) {
	rows, err := m.dbPool.Query(context.Background(), "select id,name, description, business_id, action_id, implemented ,created_at FROM risky_public.detections(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	detectionOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[DetectionModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetDetection(id string) (detectionOutput DetectionModel, err error) {
	rows, err := m.dbPool.Query(context.Background(), "select id,name, description, business_id, action_id, implemented ,created_at FROM risky_public.get_detection(fn_detection_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	detectionOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[DetectionModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteDetection(id string) (err error) {
	_, err = m.dbPool.Query(context.Background(), "select risky_public.delete_detection(fn_detection_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateDetection(detectionInput DetectionModel) (err error) {
	_, err = m.dbPool.Query(context.Background(),
		`select risky_public.create_detection(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3, 
			fn_action_id => $4, 
			fn_implemented => $5)`,
		detectionInput.Name,
		detectionInput.Description,
		detectionInput.BusinessID,
		detectionInput.ActionID,
		detectionInput.Implemented)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) UpdateDetection(detectionInput DetectionModel) (err error) {

	_, err = m.dbPool.Query(context.Background(),
		`select risky_public.update_detection(
			fn_detection_id => $1
			fn_name => $2, 
			fn_description => $3, 
			fn_business_id => $4, 
			fn_action_id => $5, 
			fn_implemented => $6)`,
		detectionInput.ID,
		detectionInput.Name,
		detectionInput.Description,
		detectionInput.BusinessID,
		detectionInput.ActionID,
		detectionInput.Implemented)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
