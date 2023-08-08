package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetDetections(businessID string) (detectionOutput []types.Detection, err error) {
	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.detection WHERE business_id = $1", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	detectionOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Detection])
	if err != nil {
		log.Printf("GetDetections Error %s", err)
		return
	}

	return
}

func (m *DBManager) GetDetection(id string) (detectionOutput types.Detection, err error) {
	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.detection WHERE id = $1", id)
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
	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_public.detection WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateDetection(detectionInput types.Detection) (detectionId string, err error) {
	err = m.DBPool.QueryRow(context.Background(),
		`INSERT INTO risky_public.detection(name, description, business_id, implemented) values($1, $2, $3, $4) RETURNING id;`,
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
		`UPDATE risky_public.detection SET name = $2, description = $3, business_id = $4,implemented = $5 WHERE id = $1;`,
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
