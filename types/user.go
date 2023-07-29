package types

type User struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	OrganizationID string `json:"organization_id" db:"organization_id"`
	GroupID        string `json:"group_id" db:"group_id"`
	CreatedAt      string `json:"created_at" db:"created_at"`
}

type Users []User
