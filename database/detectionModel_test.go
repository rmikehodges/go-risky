package database_test

import (
	"context"
	"go-risky/database"
	"testing"

	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestGetDetections(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	detections, _ := dbManager.GetDetections(businessId)

	for _, detection := range detections {
		assert.IsEqual(detection.BusinessID.String(), businessId)
	}
}

func TestGetDetection(t *testing.T) {
	var detectionId = "535705bc-fddb-4e2a-8c1c-196755ce16b6"
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	detection, _ := dbManager.GetDetection(detectionId)

	assert.IsEqual(detection.ID.String(), detectionId)
}

func TestDeleteDetection(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	detectionInput := database.DetectionModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	detectionId, _ := dbManager.CreateDetection(detectionInput)

	err = dbManager.DeleteDetection(detectionId)

	assert.Equal(t, err, nil)

	_, err = dbManager.GetDetection(detectionId)

	assert.NotEqual(t, err, nil)

}

func TestCreateDetection(t *testing.T) {
	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	detectionInput := database.DetectionModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	detectionId, err := dbManager.CreateDetection(detectionInput)

	assert.Equal(t, err, nil)

	detection, err := dbManager.GetDetection(detectionId)

	assert.Equal(t, err, nil)

	assert.Equal(t, detection.ID.String(), detectionId)
}

func TestUpdateDetection(t *testing.T) {

	poolConfig, _ := pgxpool.ParseConfig("postgres://postgres:postgres@localhost/risky")
	pgPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(err)
	}
	defer pgPool.Close()
	dbManager := &database.DBManager{DBPool: pgPool}
	createDetectionInput := database.DetectionModel{Name: "test", BusinessID: uuid.MustParse(businessId)}
	detectionId, _ := dbManager.CreateDetection(createDetectionInput)

	updateDetectionInput := createDetectionInput
	updateDetectionInput.Name = "test2"
	updateDetectionInput.ID = uuid.MustParse(detectionId)

	err = dbManager.UpdateDetection(updateDetectionInput)

	assert.Equal(t, err, nil)

	updatedDetection, _ := dbManager.GetDetection(detectionId)

	assert.Equal(t, updateDetectionInput.Name, updatedDetection.Name)
}
