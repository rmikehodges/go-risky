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

type ResourceModel struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  zeronull.Text `json:"description"`
	Cost         float32       `json:"cost" db:"cost"`
	Unit         string        `json:"unit" db:"unit"`
	Total        float32       `json:"total"`
	ResourceType string        `json:"resourceType" db:"resource_type"`
	BusinessID   uuid.UUID     `json:"businessId" db:"business_id"`
	CreatedAt    time.Time     `json:"createdAt" db:"created_at"`
}

func GetResources(businessID string) (resourceOutput []ResourceModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, cost, unit, total, resource_type, business_id, created_at FROM risky_public.resources(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	resourceOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[ResourceModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func GetResource(id string) (resourceOutput ResourceModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, cost, unit, total, resource_type, business_id, create_at FROM risky_public.get_resource(fn_resource_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	resourceOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[ResourceModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func DeleteResource(id string) (err error) {
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

	_, err = dbpool.Query(context.Background(), "select risky_public.delete_resource(fn_resource_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func CreateResource(resourceInput ResourceModel) (err error) {
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
		`select risky_public.create_resource(
			fn_name => $1, 
			fn_description => $2, 
			fn_cost=> $3, 
			fn_unit => $4, 
			fn_total => $5, 
			business_id => $6)`,
		resourceInput.Name,
		resourceInput.Description,
		resourceInput.Cost,
		resourceInput.Unit,
		resourceInput.Total,
		resourceInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func UpdateResource(resourceInput ResourceModel) (err error) {
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
		`select risky_public.update_resource(
			fn_resource_id => $1
			fn_name => $2, 
			fn_description => $3, 
			fn_cost=> $4, 
			fn_unit => $5, 
			fn_total => $6, 
			business_id => $7)`,
		resourceInput.ID,
		resourceInput.Name,
		resourceInput.Description,
		resourceInput.Cost,
		resourceInput.Unit,
		resourceInput.Total,
		resourceInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
