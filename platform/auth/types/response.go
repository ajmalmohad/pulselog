package types

type SuccessResponse struct {
	Message string        `json:"message"`
	Data    TokenResponse `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
