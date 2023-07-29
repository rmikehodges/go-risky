package database

import (
	"context"
	"log"

	"go-risky/types"

	"github.com/jackc/pgx/v5"
)

func (m *DBManager) GetOrganizations() (organizationOutput []types.Organization, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.organization")
	if err != nil {
		log.Println(err)
		return
	}

	organizationOutput, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Organization])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetOrganization(id string) (organizationOutput types.Organization, err error) {

	rows, err := m.DBPool.Query(context.Background(), "SELECT * FROM risky_public.organization WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	organizationOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Organization])
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) GetOrganizationByUserEmail(email string) (organizationOutput types.Organization, err error) {
	rows, err := m.DBPool.Query(context.Background(), `SELECT
	risky_secret.organization.id,
	risky_secret.organization.name,
	risky_secret.organization.oauth_client_id,
	risky_secret.organization.oauth_client_secret,
	risky_secret.organization.redirect_uri,
	risky_secret.organization.scopes,
	risky_secret.organization.oauth_endpoint,
	risky_secret.organization.oauth_enabled,
	risky_secret.organization.created_at
	FROM risky_secret.users 
	INNER JOIN risky_secret.organization ON risky_secret.users.organization_id = risky_secret.organization.id 
	WHERE risky_secret.users.email = $1`, email)
	if err != nil {
		log.Println(err)
		return
	}

	organizationOutput, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Organization])
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (m *DBManager) DeleteOrganization(id string) (err error) {
	_, err = m.DBPool.Exec(context.Background(), "DELETE FROM risky_public.organization WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func (m *DBManager) CreateOrganization(organizationInput types.Organization) (createdOrganization string, err error) {
	err = m.DBPool.QueryRow(context.Background(),
		`INSERT INTO risky_public.organization(
			name, 
			oauth_client_id,
			oauth_client_secret,
			redirect_uri,
			scopes,
			oauth_endpoint) 
			 values($1, $2, $3, $4, $5, $6) 
			 RETURNING id`,
		organizationInput.Name,
		organizationInput.OAuthClientID,
		organizationInput.OAuthClientSecret,
		organizationInput.RedirectURI,
		organizationInput.Scopes,
		organizationInput.OAuthEndpoint).Scan(&createdOrganization)

	if err != nil {
		log.Printf("Error creating organization %s", err)
		return
	}
	return
}

func (m *DBManager) UpdateOrganization(organizationInput types.Organization) (err error) {
	_, err = m.DBPool.Exec(context.Background(),
		`UPDATE risky_public.organization SET 
		name = $2, 
		oauth_client_id = $3,
		oauth_client_secret = $4,
		redirect_uri = $5,
		scopes = $6,
		oauth_endpoint = $7
		WHERE id = $1`,
		organizationInput.ID,
		organizationInput.Name,
		organizationInput.OAuthClientID,
		organizationInput.OAuthClientSecret,
		organizationInput.RedirectURI,
		organizationInput.Scopes,
		organizationInput.OAuthEndpoint)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
