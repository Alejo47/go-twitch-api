package authorization

type Scopes []string

type AccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
	Scope        Scopes `json:"scope"`
	TokenType    string `json:"token_type"`
	StatusCode   uint
}

type RefreshAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type ValidateTokenResponse struct {
	ClientID  string `json:"client_id"`
	Login     string `json:"login"`
	Scopes    Scopes `json:"scopes"`
	UserID    string `json:"user_id"`
	ExpiresIn int    `json:"expires_in"`
}
