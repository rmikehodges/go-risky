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

type CapabilityModel struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at"`
}

func GetCapabilities(businessID string) (capabilityOutput []CapabilityModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, capability_id, vulnerability_id, business_id, complexity, asset_id, created_at FROM risky_public.capabilities(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	capabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[CapabilityModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func GetCapability(id string) (capabilityOutput CapabilityModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, capability_id, vulnerability_id, business_id, complexity, asset_id, created_at FROM risky_public.get_capability(fn_capability_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	capabilityOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[CapabilityModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func DeleteCapability(id string) (err error) {
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

	_, err = dbpool.Query(context.Background(), "select risky_public.delete_capability(fn_capability_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func CreateCapability(capabilityInput CapabilityModel) (err error) {
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
		`select risky_public.create_capability(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3)`,
		capabilityInput.Name,
		capabilityInput.Description,
		capabilityInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func UpdateCapability(capabilityInput CapabilityModel) (err error) {
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
		`select risky_public.update_capability(
			fn_capability_id => $1
			fn_name => $2, 
			fn_description => $3,
			fn_business_id => $4)`,
		capabilityInput.ID,
		capabilityInput.Name,
		capabilityInput.Description,
		capabilityInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
