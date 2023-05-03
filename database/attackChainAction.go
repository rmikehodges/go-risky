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
	"github.com/jackc/pgx/v5/pgxpool"
)

// Create type AttackChainActionModel based on table defintition risky_public.attack_chain_action in database/tables.sql
type AttackChainActionModel struct {
	BusinessID    uuid.UUID `db:"business_id"`
	ActionID      uuid.UUID `db:"action_id"`
	AttackChainID uuid.UUID `db:"attack_chain_id"`
	Position      int       `db:"position"`
	CreatedAt     time.Time `db:"created_at"`
}

func GetAttackChainActions(businessID string) (actionOutput []AttackChainActionModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select attack_chain_id, action_id,position,business_id, created_at FROM risky_public.attack_chain_actions(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	actionOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[AttackChainActionModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func GetAttackChainAction(id string) (actionOutput AttackChainActionModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select  attack_chain_id, action_id,position,business_id, FROM risky_public.get_attack_chain_action(fn_attack_chain_action_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	actionOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[AttackChainActionModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func DeleteAttackChainAction(id string) (err error) {
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

	_, err = dbpool.Query(context.Background(), "select risky_public.delete_attack_chain_action(fn_attack_chain_action_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func CreateAttackChainAction(attackChainActionInput AttackChainActionModel) (err error) {
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
		`select risky_public.create_attack_chain_action(
			fn_attack_chain_id => $1, 
			fn_action_id => $2, 
			fn_position => $3
			fn_business_id => $4)`,
		attackChainActionInput.AttackChainID,
		attackChainActionInput.ActionID,
		attackChainActionInput.Position,
		attackChainActionInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func UpdateAttackChainAction(attackChainActionInput AttackChainActionModel) (err error) {
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
		`select risky_public.update_attack_chain_action(
			fn_attack_chain_id => $1, 
			fn_action_id => $2, 
			fn_position => $3
			fn_business_id => $4)`,
		attackChainActionInput.AttackChainID,
		attackChainActionInput.ActionID,
		attackChainActionInput.Position,
		attackChainActionInput.BusinessID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
