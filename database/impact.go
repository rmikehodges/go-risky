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

type ImpactModel struct {
	ID               uuid.UUID       `json:"id"`
	Name             string          `json:"name"`
	Description      zeronull.Text   `json:"description"`
	BusinessID       uuid.UUID       `json:"businessId" db:"business_id"`
	ThreatID         uuid.UUID       `json:"threatId" db:"threat_id"`
	ExploitationCost zeronull.Float8 `json:"exploitationCost" db:"exploitation_cost"`
	MitigationCost   zeronull.Float8 `json:"mitigationCost" db:"mitigation_cost"`
	CreatedAt        time.Time       `json:"createdAt" db:"created_at"`
}

func GetImpacts(businessID string) (impactOutput []ImpactModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, business_id, threat_id, exploitation_cost, mitigation_cost, created_at FROM risky_public.impacts(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	impactOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[ImpactModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func GetImpact(id string) (impactOutput ImpactModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, business_id, threat_id, exploitation_cost, mitigation_cost, created_at FROM risky_public.get_impact(fn_impact_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	impactOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[ImpactModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func DeleteImpact(id string) (err error) {
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

	_, err = dbpool.Query(context.Background(), "select risky_public.delete_impact(fn_impact_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func CreateImpact(impactInput ImpactModel) (err error) {
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
		`select risky_public.create_impact(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3, 
			fn_threat_id => $4)`,
		impactInput.Name,
		impactInput.Description,
		impactInput.BusinessID,
		impactInput.ThreatID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func UpdateImpact(impactInput ImpactModel) (err error) {
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
		`select risky_public.update_impact(
			fn_impact_id => $1
			fn_name => $2, 
			fn_description => $3, 
			fn_business_id => $4, 
			fn_threat_id => $5)`,
		impactInput.ID,
		impactInput.Name,
		impactInput.Description,
		impactInput.BusinessID,
		impactInput.ThreatID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
