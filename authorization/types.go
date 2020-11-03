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

type ResponseError *struct {
	InternalError error
	Error         string `json:"error"`
	Status        int    `json:"status"`
	Message       string `json:"message"`
}
