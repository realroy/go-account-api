package models

import (
	"errors"
	"fmt"
	"go-account-api/schemas"
	"reflect"
	"strconv"
)

var id = 1
var db = make(map[int]*schemas.Account)

// Account Model
type Account struct{}

// Create accounts
func (Account) Create(accounts ...*schemas.CreateAccountArg) ([]*schemas.Account, error) {
	newAccounts := make([]*schemas.Account, 0)
	for _, account := range accounts {
		db[id] = &schemas.Account{ID: strconv.Itoa(id), Email: account.Email, Password: account.Password}
		newAccounts = append(newAccounts, db[id])
		id++
	}
	return newAccounts, nil
}

// All accounts
func (Account) All() ([]*schemas.Account, error) {
	accounts := make([]*schemas.Account, 0)
	for _, v := range db {
		accounts = append(accounts, v)
	}
	return accounts, nil
}

// FindByID of account
func (Account) FindByID(id string) (*schemas.Account, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	a, ok := db[parsedID]
	if !ok {
		return nil, errors.New("Account is not found")
	}
	return a, nil
}

// Find accounts with struct's field name and value
func (Account) Find(fieldName string, value string) ([]*schemas.Account, error) {
	a := &schemas.Account{}
	elem := reflect.ValueOf(a).Elem()
	if !elem.FieldByName(fieldName).IsValid() {
		msg := fmt.Sprintf("%s is not exist in Account struct", fieldName)
		return nil, errors.New(msg)
	}

	matches := make([]*schemas.Account, 0)
	for _, account := range db {
		elem = reflect.ValueOf(account).Elem()
		if elem.FieldByName(fieldName).Interface() == value {
			matches = append(matches, account)
		}
	}
	if len(matches) == 0 {
		return nil, errors.New("Account is not found")
	}
	return matches, nil
}

// Update an exist account
func (Account) Update(id string, new *schemas.CreateAccountArg) (*schemas.Account, error) {
	parsedID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	old, found := db[parsedID]
	if !found {
		return nil, errors.New("Account is not found")
	}

	db[parsedID] = &schemas.Account{ID: old.ID, Email: new.Email, Password: new.Password}
	return db[parsedID], nil
}

// Destroy an account by email
func (Account) Destroy(ids ...int) error {
	for _, id := range ids {
		delete(db, id)
	}
	return nil
}

// DestroyAll accounts
func (Account) DestroyAll() error {
	for k := range db {
		delete(db, k)
	}
	return nil
}

// Count of account
func (Account) Count() int {
	return len(db)
}
