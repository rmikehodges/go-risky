package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetDetections(businessID string) (detectionOutput []types.Detection, err error) {
	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, implemented ,created_at FROM risky_public.detections(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	detectionOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Detection])
	if err != nil {
		log.Printf("GetDetections ErrorL %s", err)
		return
	}

	return
}

func (m *DBManager) GetDetection(id string) (detectionOutput types.Detection, err error) {
	rows, err := m.DBPool.Query(context.Background(), "select id,name, description, business_id, implemented ,created_at FROM risky_public.get_detection(fn_detection_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	detectionOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Detection])
	if err != nil {
		log.Printf("GetDetection Error: %s", err)
		return
	}

	return
}

func (m *DBManager) DeleteDetection(id string) (err error) {
	_, err = m.DBPool.Exec(context.Background(), "select risky_public.delete_detection(fn_detection_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateDetection(detectionInput types.Detection) (detectionId string, err error) {
	err = m.DBPool.QueryRow(context.Background(),
		`select risky_public.create_detection(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3, 
			fn_implemented => $4)`,
		detectionInput.Name,
		detectionInput.Description,
		detectionInput.BusinessID,
		detectionInput.Implemented).Scan(&detectionId)
	if err != nil {
		log.Printf("CreateDetection Error: %s", err)
		return
	}

	return
}

func (m *DBManager) UpdateDetection(detectionInput types.Detection) (err error) {

	_, err = m.DBPool.Exec(context.Background(),
		`select risky_public.update_detection(
			fn_detection_id => $1,
			fn_name => $2, 
			fn_description => $3, 
			fn_business_id => $4, 
			fn_implemented => $6)`,
		detectionInput.ID,
		detectionInput.Name,
		detectionInput.Description,
		detectionInput.BusinessID,
		detectionInput.Implemented)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
