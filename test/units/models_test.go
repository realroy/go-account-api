package units

import (
	"go-account-api/models"
	"go-account-api/schemas"
	"testing"

	"github.com/stretchr/testify/assert"
)

var m = models.Account{}

func AfterEach() {
	m.DestroyAll()
}

func MatchAccount(t *testing.T, actual *schemas.Account, expect *schemas.Account) {
	assert.Equal(t, actual.Password, expect.Password)
	assert.Equal(t, actual.Email, expect.Email)
}

func TestAccountModelCreate(t *testing.T) {
	defer AfterEach()

	prevCount := m.Count()
	a := &schemas.CreateAccountArg{Email: "a@a.com", Password: "123456789"}

	_, err := m.Create(a)
	assert.Nil(t, err)
	assert.Equal(t, m.Count(), (prevCount + 1))
}

func TestAccountModelAll(t *testing.T) {
	defer AfterEach()

	expects := []*schemas.CreateAccountArg{
		&schemas.CreateAccountArg{Email: "a@a.com", Password: "123456789"},
		&schemas.CreateAccountArg{Email: "b@b.com", Password: "123456789"},
	}

	for _, e := range expects {
		m.Create(e)
	}

	actuals, _ := m.All()
	for i := range expects {
		actual, expect := actuals[i], expects[i]

		assert.Equal(t, actual.Password, expect.Password)
		assert.Equal(t, actual.Email, expect.Email)
	}
}

func TestAccountModelFindByID(t *testing.T) {
	defer AfterEach()

	a := &schemas.CreateAccountArg{Email: "a@a.com", Password: "123456789"}
	expects, _ := m.Create(a)
	expect := expects[0]
	actual, err := m.FindByID(expect.ID)
	assert.Nil(t, err)

	MatchAccount(t, actual, expect)
}

func TestAccountModelFind(t *testing.T) {
	a := &schemas.CreateAccountArg{Email: "a@a.com", Password: "123456789"}
	expects, _ := m.Create(a)
	actuals, err := m.Find("Email", expects[0].Email)
	assert.Nil(t, err)
	MatchAccount(t, actuals[0], expects[0])
}

func TestAccountModelUpdate(t *testing.T) {
	new := &schemas.CreateAccountArg{Email: "b@b.com", Password: "123456780"}

	olds, _ := m.Create(&schemas.CreateAccountArg{Email: "a@a.com", Password: "123456789"})
	expect, err := m.Update(olds[0].ID, new)

	assert.Nil(t, err)

	actual, _ := m.FindByID(olds[0].ID)
	MatchAccount(t, actual, expect)
}
