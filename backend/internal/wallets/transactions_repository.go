package wallets

import (
	"github.com/renanmedina/investment-warlock/utils"
)

func SaveTransaction(transaction Transaction) Transaction {
	db := utils.GetDatabase()

	// try creating if fails probably exists then updated (should improve this in the future)
	if _, errCreate := db.Create("transactions", transaction.ToMap()); errCreate != nil {
		if _, errUpdate := db.Change(transaction.Id, transaction.ToMap()); errUpdate != nil {
			panic(errUpdate)
		}
	}

	return transaction
}
