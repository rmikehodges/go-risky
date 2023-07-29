package types

type Group struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	OrganizationID string `json:"organization_id"`
	CreatedAt      string `json:"created_at"`
}

type Groups []Group
