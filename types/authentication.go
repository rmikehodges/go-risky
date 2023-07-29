package types

type Authentication struct {
	ID                 string `json:"id"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	PasswordResetToken string `json:"password_reset_token" db:"password_reset_token"`
}
