package types

type Organization struct {
	ID                string  `json:"id"`
	Name              string  `json:"name"`
	OAuthClientID     *string `json:"oauth_client_id"`
	OAuthClientSecret *string `json:"oauth_client_secret"`
	OAuthEnabled      bool    `json:"oauth_enabled"`
	RedirectURI       *string `json:"redirect_uri"`
	Scopes            *string `json:"scopes"`
	OAuthEndpoint     *string `json:"ouath_endpoint"`
	CreatedAt         string  `json:"created_at"`
}
