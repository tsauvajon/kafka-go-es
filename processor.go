package main

// Process : processes a CreateEvent
func (e CreateEvent) Process() (*BankAccount, error) {
	return updateAccount(e.AccID, map[string]interface{}{
		"ID":      e.AccID,
		"Name":    e.AccName,
		"Balance": "0",
	})
}

// Process : processes a DepositEvent
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

// // Process :
// func (e InvalidEvent) Process() error {
// 	return nil
// }

// // Process :
// func (e AcceptEvent) Process() error {
// 	return nil
// }
