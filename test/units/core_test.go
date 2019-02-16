package units

import (
	"go-account-api/core"
	"go-account-api/schemas"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	req := &schemas.LoginRequest{
		Email:    "a@a.com",
		Password: "123456789",
	}
	_, status := core.Login(req)
	assert.Equal(t, http.StatusOK, status)
}

func TestLogoutSuccess(t *testing.T) {
	req := &schemas.LogoutRequest{
		RefreshToken: "123456789",
	}
	_, status := core.Logout(req)
	assert.Equal(t, http.StatusNoContent, status)
}
func TestLogoutWithEmptyRefreshToken(t *testing.T) {
	req := &schemas.LogoutRequest{
		RefreshToken: "",
	}
	_, status := core.Logout(req)
	assert.Equal(t, http.StatusBadRequest, status)
}
func TestRegister(t *testing.T) {
	req := &schemas.RegisterRequest{
		Email:    "a@a.com",
		Password: "123456789",
	}
	_, status := core.Register(req)
	assert.Equal(t, http.StatusOK, status)
}
