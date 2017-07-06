package main

import (
	"errors"
)

// Process : processes a CreateEvent => creates a bank account
func (e CreateEvent) Process() (*BankAccount, error) {
	return updateAccount(e.AccID, map[string]interface{}{
		"ID":      e.AccID,
		"Name":    e.AccName,
		"Balance": "0",
	})
}

// Process : processes a DepositEvent => adds money to an account
func (e DepositEvent) Process() (*BankAccount, error) {
	acc, err := FetchAccount(e.AccID)
	if err != nil {
		return nil, err
	}

	newBalance := acc.Balance + e.Amount
	return updateAccount(e.AccID, map[string]interface{}{
		"Balance": newBalance,
	})
}

// Process : processes a WithdrawEvent => remove money from an account
func (e WithdrawEvent) Process() (*BankAccount, error) {
	acc, err := FetchAccount(e.AccID)
	if err != nil {
		return nil, err
	}

	newBalance := acc.Balance - e.Amount
	return updateAccount(e.AccID, map[string]interface{}{
		"Balance": newBalance,
	})
}

// Process : processes a TransferEvent => transfers money from an account to another
func (e TransferEvent) Process() (*BankAccount, error) {
	acc, err := FetchAccount(e.AccID)

	if err != nil {
		return nil, err
	}

	if acc.Balance < e.Amount {
		return nil, errors.New("Insufficient balance")
	}

	target, err := FetchAccount(e.TargetID)

	if err != nil {
		return nil, err
	}

	newBalance := acc.Balance - e.Amount
	newTargetBalance := target.Balance + e.Amount

	target, err = updateAccount(
		e.TargetID, map[string]interface{}{
			"Balance": newTargetBalance,
		})

	return updateAccount(e.AccID, map[string]interface{}{
		"Balance": newBalance,
	})
}

// // Process :
// func (e InvalidEvent) Process() error {
// 	return nil
// }

// // Process :
// func (e AcceptEvent) Process() error {
// 	return nil
// }
