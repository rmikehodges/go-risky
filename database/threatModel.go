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

type ThreatModel struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at"`
}

func GetThreats(businessID string) (threatOutput []ThreatModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, business_id,created_at FROM risky_public.threats(fn_business_id => $1)", businessID)
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

func GetThreat(id string) (threatOutput ThreatModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description,business_id, created_at FROM risky_public.get_threat(fn_threat_id => $1)", id)
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

func DeleteThreat(id string) (err error) {
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

	_, err = dbpool.Query(context.Background(), "select risky_public.delete_threat(fn_threat_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func CreateThreat(threatInput ThreatModel) (err error) {
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

func UpdateThreat(threatInput ThreatModel) (err error) {
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
