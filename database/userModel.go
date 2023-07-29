package database

import (
	"context"
	"fmt"
	"log"

	"go-risky/types"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetUsers(businessID string) (userOutput []types.User, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_secret.user WHERE organizationId = $1", businessID)
	if err != nil {
		log.Println(err)
		return
	}

	userOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.User])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetUser(id string) (userOutput types.User, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_secret.user WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(rows)

	userOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.User])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetUserByEmail(email string) (userOutput types.User, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_secret.user WHERE email = $1", email)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(rows)

	userOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.User])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) DeleteUser(id string) (err error) {
	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_secret.user WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateUser(userInput types.User) (createdUser string, err error) {
	err = m.DBPool.QueryRow(context.Background(),
		`INSERT INTO risky_secret.user(
			email, 
			group_id, 
			organization_id) 
			 values($1, $2, $3) 
			 RETURNING id`,
		userInput.Email,
		userInput.GroupID,
		userInput.OrganizationID).Scan(&createdUser)

	if err != nil {
		log.Printf("Error creating user %s", err)
		return
	}
	return
}

func (m *DBManager) UpdateUser(userInput types.User) (err error) {
	_, err = m.DBPool.Exec(context.Background(),
		`UPDATE risky_secret.user SET 
		email = $2, 
		group_id = $3, 
		organization_id = $4
		WHERE id = $1`,
		userInput.ID,
		userInput.Email,
		userInput.GroupID,
		userInput.OrganizationID)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (m *DBManager) GetUsersByGroup(groupID string) (userOutput []types.User, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_secret.user WHERE group_id = $1", groupID)
	if err != nil {
		log.Println(err)
		return
	}

	userOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.User])
	if err != nil {
		log.Println(err)
		return
	}

	return
}
