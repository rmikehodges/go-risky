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

type BusinessModel struct {
	ID        uuid.UUID       `json:"id"`
	Name      string          `json:"name"`
	Revenue   zeronull.Float8 `json:"revenue"`
	CreatedAt time.Time       `json:"createdAt" db:"created_at"`
}

func GetBusinesss(businessID string) (businessOutput []BusinessModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, revenue, capability_id, created_at FROM risky_public.businesses(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	businessOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[BusinessModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func GetBusiness(id string) (businessOutput BusinessModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, revenue, created_at FROM risky_public.get_business(fn_business_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	businessOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[BusinessModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func DeleteBusiness(id string) (err error) {
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

	_, err = dbpool.Query(context.Background(), "select risky_public.delete_business(fn_business_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func CreateBusiness(businessInput BusinessModel) (err error) {
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
		`select risky_public.create_business(
			fn_name => $1, 
			fn_revenue => $2)`,
		businessInput.Name,
		businessInput.Revenue)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func UpdateBusiness(businessInput BusinessModel) (err error) {
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
		`select risky_public.update_business(
			fn_business_id => $1
			fn_name => $2, 
			fn_description => $3)`,
		businessInput.ID,
		businessInput.Name,
		businessInput.Revenue)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
