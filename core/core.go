package core

import (
	"go-account-api/schemas"
	"net/http"
)

func Login(req *schemas.LoginRequest) (*schemas.LoginResponse, int) {
	res := &schemas.LoginResponse{
		AccessToken:  "0ddb922452c983a70566e30dce16e2017db335103e35d783874c448862a78168",
		TokenType:    "bearer",
		ExpiresIn:    7200,
		RefreshToken: "f2188c4165d912524e04c6496d10f06803cc08ed50271a0b0a73061e3ac1c06c",
	}
	return res, http.StatusOK
}

func Logout(req *schemas.LogoutRequest) (*interface{}, int) {
	if len(req.RefreshToken) == 0 {
		return nil, http.StatusBadRequest
	}
	return nil, http.StatusNoContent
}

func Register(req *schemas.RegisterRequest) (*schemas.RegisterResponse, int) {
	res := &schemas.RegisterResponse{
		ID:    1,
		Email: req.Email,
	}
	return res, http.StatusOK
}
