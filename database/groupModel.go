package database

import (
	"context"
	"fmt"
	"log"

	"go-risky/types"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetGroups(organizationId string) (groupOutput []types.Group, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.group WHERE organization_id = $1", organizationId)
	if err != nil {
		log.Println(err)
		return
	}

	groupOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Group])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetGroup(id string) (groupOutput types.Group, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.group WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(rows)

	groupOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Group])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteGroup(id string) (err error) {
	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_public.group WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateGroup(groupInput types.Group) (createdGroup string, err error) {
	err = m.DBPool.QueryRow(context.Background(),
		`INSERT INTO risky_public.group(
			name, 
			organization_id) 
			 values($1, $2) 
			 RETURNING id`,
		groupInput.Name,
		groupInput.OrganizationID).Scan(&createdGroup)

	if err != nil {
		log.Printf("Error creating group %s", err)
		return
	}
	return
}

func (m *DBManager) UpdateGroup(groupInput types.Group) (err error) {
	_, err = m.DBPool.Exec(context.Background(),
		`UPDATE risky_public.group SET 
		name = $2, 
		organization_id = $3 
		WHERE id = $1`,
		groupInput.ID,
		groupInput.Name,
		groupInput.OrganizationID)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
