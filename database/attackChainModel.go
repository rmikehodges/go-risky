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

type AttackChainModel struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Description zeronull.Text `json:"description"`
	BusinessID  uuid.UUID     `json:"businessId" db:"business_id"`
	ThreatID    uuid.UUID     `json:"assetId" db:"asset_id"`
	CreatedAt   time.Time     `json:"createdAt" db:"created_at"`
}

func GetAttackChains(businessID string) (attackChainOutput []AttackChainModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, capability_id, vulnerability_id, business_id, complexity, asset_id, created_at FROM risky_public.attackChains(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	attackChainOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[AttackChainModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func GetAttackChain(id string) (attackChainOutput AttackChainModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, capability_id, vulnerability_id, business_id, complexity, asset_id, created_at FROM risky_public.get_attack_chain(fn_attack_chain_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	attackChainOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[AttackChainModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func DeleteAttackChain(id string) (err error) {
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

	_, err = dbpool.Query(context.Background(), "select risky_public.delete_attack_chain(fn_attack_chain_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func CreateAttackChain(attackChainInput AttackChainModel) (err error) {
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
		`select risky_public.create_attack_chain(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3,  
			fn_threat_id => $4)`,
		attackChainInput.Name,
		attackChainInput.Description,
		attackChainInput.BusinessID,
		attackChainInput.ThreatID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func UpdateAttackChain(attackChainInput AttackChainModel) (err error) {
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
		`select risky_public.update_attack_chain(
			fn_attack_chain_id => $1
			fn_name => $2, 
			fn_description => $3, 
			fn_business_id => $4, 
			fn_threat_id => $5)`,
		attackChainInput.ID,
		attackChainInput.Name,
		attackChainInput.Description,
		attackChainInput.BusinessID,
		attackChainInput.ThreatID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
