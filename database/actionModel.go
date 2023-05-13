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

type ActionModel struct {
	ID              uuid.UUID     `json:"id"`
	Name            string        `json:"name"`
	Description     zeronull.Text `json:"description"`
	CapabilityID    *uuid.UUID    `json:"capabilityId" db:"capability_id"`
	VulnerabilityID *uuid.UUID    `json:"vulnerabilityId" db:"vulnerability_id"`
	BusinessID      uuid.UUID     `json:"businessId" db:"business_id"`
	Complexity      zeronull.Text `json:"complexity"`
	AssetID         *uuid.UUID    `json:"assetId" db:"asset_id"`
	CreatedAt       time.Time     `json:"createdAt" db:"created_at"`
}

func GetActions(businessID string) (actionOutput []ActionModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, capability_id, vulnerability_id, business_id, complexity, asset_id, created_at FROM risky_public.actions(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	actionOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[ActionModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func GetAction(id string) (actionOutput ActionModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, capability_id, vulnerability_id, business_id, complexity, asset_id, created_at FROM risky_public.get_action(fn_action_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	actionOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[ActionModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func DeleteAction(id string) (err error) {
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

	_, err = dbpool.Query(context.Background(), "select risky_public.delete_action(fn_action_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func CreateAction(actionInput ActionModel) (err error) {
	databaseURL := os.Getenv("DATABASE_URL")

	dbconfig, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Println(err)
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
	fmt.Println("starting Create action")
	fmt.Println(actionInput.CapabilityID)
	_, err = dbpool.Exec(context.Background(),
		`select risky_public.create_action(
			fn_name => $1, 
			fn_description => $2, 
			fn_capability_id => $3, 
			fn_vulnerability_id => $4, 
			fn_business_id => $5, 
			fn_complexity => $6, 
			fn_asset_id => $7)`,
		actionInput.Name,
		actionInput.Description,
		actionInput.CapabilityID,
		actionInput.VulnerabilityID,
		actionInput.BusinessID,
		actionInput.Complexity,
		actionInput.AssetID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func UpdateAction(actionInput ActionModel) (err error) {
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
	_, err = dbpool.Exec(context.Background(),
		`select risky_public.update_action(
			fn_action_id => $1,
			fn_name => $2, 
			fn_description => $3, 
			fn_capability_id => $4, 
			fn_vulnerability_id => $5, 
			fn_business_id => $6, 
			fn_complexity => $7, 
			fn_asset_id => $8)`,
		actionInput.ID,
		actionInput.Name,
		actionInput.Description,
		actionInput.CapabilityID,
		actionInput.VulnerabilityID,
		actionInput.BusinessID,
		actionInput.Complexity,
		actionInput.AssetID)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
