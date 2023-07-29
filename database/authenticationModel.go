package database

import (
	"context"
	"go-risky/types"
	"log"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) AuthenticateWithPassword(email string, password string) (userOutput types.User, err error) {
	row, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_secret.user WHERE email = $1 AND password = $2", email, password)
	if err != nil {
		log.Println(err)
		return
	}

	userOutput, err = pgx.CollectOneRow(row, pgx.RowToStructByName[types.User])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) ChangePassword(userId string, password string) (err error) {
	_, err = m.DBPool.Exec(context.Background(), "UPDATE risky_secret.user SET password = $2, password_reset_token = NULL WHERE id = $1", userId, password)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) ResetPassword(userId string, password string, passwordResetToken string) (err error) {
	_, err = m.DBPool.Exec(context.Background(), "UPDATE risky_secret.user SET password = $2, password_reset_token = NULL WHERE id = $1 AND password_reset_token = $3", userId, password, passwordResetToken)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) InitiatePasswordReset(userId string, resetToken string) (err error) {
	_, err = m.DBPool.Exec(context.Background(), "UPDATE risky_secret.user SET password_reset_token = $2 WHERE id = $1", userId, resetToken)
	if err != nil {
		log.Printf("Initiate Password Reset Database Error: %s", err)
		return
	}

	return
}
