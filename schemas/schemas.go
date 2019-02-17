package schemas

type (
	CreateAccountArg struct {
		Email    string `json: "email"`
		Password string `json: "password`
	}
	Account struct {
		ID       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
	}
	LogoutRequest struct {
		RefreshToken string `json:"refresh_token"`
	}
	RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RegisterResponse struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
	}
)
