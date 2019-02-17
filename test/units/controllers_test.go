package units

import (
	"go-account-api/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestLoginWithValidAccount(t *testing.T) {
	e := echo.New()
	account := `{ "email": "r@s.com", "password": "123456789" }`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(account))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
func TestLoginWithInvalidAccount(t *testing.T) {
	e := echo.New()
	account := `{ "email": "stranger@stranger.com", "password": "123456789" }`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(account))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, controllers.Login(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}
