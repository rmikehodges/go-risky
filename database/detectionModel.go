package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
	"github.com/jackc/pgx/v5/pgxpool"
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

func GetDetections(businessID string) (detectionOutput []DetectionModel, err error) {
	databaseURL := os.Getenv("DATABASE_URL")

	dbconfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return
	}
	dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), dbconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		return
	}
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), "select id,name, description, business_id, action_id, implemented ,created_at FROM risky_public.detections(fn_business_id => $1)", businessID)
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

func GetDetection(id string) (detectionOutput DetectionModel, err error) {
	databaseURL := os.Getenv("DATABASE_URL")

	dbconfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return
	}
	dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), dbconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		return
	}
	defer dbpool.Close()

	rows, err := dbpool.Query(context.Background(), "select id,name, description, business_id, action_id, implemented ,created_at FROM risky_public.get_detection(fn_detection_id => $1)", id)
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

func DeleteDetection(id string) (err error) {
	databaseURL := os.Getenv("DATABASE_URL")

	dbconfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return
	}
	dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), dbconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		return
	}
	defer dbpool.Close()

	_, err = dbpool.Query(context.Background(), "select risky_public.delete_detection(fn_detection_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func CreateDetection(detectionInput DetectionModel) (err error) {
	databaseURL := os.Getenv("DATABASE_URL")

	dbconfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return
	}
	dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), dbconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		return
	}
	defer dbpool.Close()

	_, err = dbpool.Query(context.Background(),
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

func UpdateDetection(detectionInput DetectionModel) (err error) {
	databaseURL := os.Getenv("DATABASE_URL")

	dbconfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return
	}
	dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), dbconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		return
	}
	defer dbpool.Close()

	_, err = dbpool.Query(context.Background(),
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
