package main

import (
	"errors"
	"strconv"
)

// BankAccount : bank account
type BankAccount struct {
	ID      string
	Name    string
	Balance int
}

// FetchAccount : Fetches the account
func FetchAccount(id string) (*BankAccount, error) {
	cmd := Redis.HGetAll(id)
	if err := cmd.Err(); err != nil {
		return nil, err
	}

	data := cmd.Val()
	if len(data) == 0 {
		return nil, nil
	}

	return ToAccount(data)
}

func updateAccount(id string, data map[string]interface{}) (*BankAccount, error) {
	cmd := Redis.HMSet(id, data)

	if err := cmd.Err(); err != nil {
		return nil, err
	}

	return FetchAccount(id)
}

// ToAccount : ??
func ToAccount(m map[string]string) (*BankAccount, error) {
	balance, err := strconv.Atoi(m["Balance"])
	if err != nil {
		return nil, err
	}

	if _, ok := m["ID"]; !ok {
		return nil, errors.New("Missing account ID")
	}

	return &BankAccount{
		ID:      m["ID"],
		Name:    m["Name"],
		Balance: balance,
	}, nil
}
