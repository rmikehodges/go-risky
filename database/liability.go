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

type LiabilityModel struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	Description  zeronull.Text `json:"description"`
	Quantity     float32       `json:"quantity"`
	Cost         float32       `json:"cost"`
	BusinessID   uuid.UUID     `json:"businessId" db:"business_id"`
	MitigationID uuid.UUID     `json:"mitigationId" db:"mitigation_id"`
	ResourceID   uuid.UUID     `json:"resourceId" db:"resource_id"`
	ThreatID     uuid.UUID     `json:"threatId" db:"threat_id"`
	ImpactID     uuid.UUID     `json:"impactId" db:"impact_id"`
	CreatedAt    time.Time     `json:"createdAt" db:"created_at"`
}

func GetLiabilities(businessID string) (liabilityOutput []LiabilityModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, business_id, mitigation_id, resource_id, threat_id, impact_id, created_at FROM risky_public.liabilities(fn_business_id => $1)", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	liabilityOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[LiabilityModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func GetLiability(id string) (liabilityOutput LiabilityModel, err error) {
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

	rows, err := dbpool.Query(context.Background(), "select id,name, description, business_id, mitigation_id, resource_id, threat_id, impact_id, created_at FROM risky_public.get_liability(fn_liability_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	liabilityOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[LiabilityModel])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func DeleteLiability(id string) (err error) {
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

	_, err = dbpool.Query(context.Background(), "select risky_public.delete_liability(fn_liability_id => $1)", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func CreateLiability(liabilityInput LiabilityModel) (err error) {
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
		`select risky_public.create_liability(
			fn_name => $1, 
			fn_description => $2, 
			fn_business_id => $3, 
			fn_mitigation_id => $4, 
			fn_resource_id => $5, 
			fn_threat_id => $6, 
			fn_impact_id => $7)`,
		liabilityInput.Name,
		liabilityInput.Description,
		liabilityInput.BusinessID,
		liabilityInput.MitigationID,
		liabilityInput.ResourceID,
		liabilityInput.ThreatID,
		liabilityInput.ImpactID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func UpdateLiability(liabilityInput LiabilityModel) (err error) {
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
		`select risky_public.update_liability(
			fn_liability_id => $1
			fn_name => $2, 
			fn_description => $3, 
			fn_business_id => $4, 
			fn_mitigation_id => $5, 
			fn_resource_id => $6, 
			fn_threat_id => $7, 
			fn_impact_id => $8)`,
		liabilityInput.ID,
		liabilityInput.Name,
		liabilityInput.Description,
		liabilityInput.BusinessID,
		liabilityInput.MitigationID,
		liabilityInput.ResourceID,
		liabilityInput.ThreatID,
		liabilityInput.ImpactID)
	if err != nil {
		log.Println(err)
		return
	}

	return
}
