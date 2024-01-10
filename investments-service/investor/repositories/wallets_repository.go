package repositories

import (
	"github.com/renanmedina/investment-warlock/investments-service/investor"
	"github.com/renanmedina/investment-warlock/investments-service/utils"
)

func Save(wallet investor.Wallet) investor.Wallet {
	db := utils.GetDatabase()

	// try creating if fails probably exists then updated (should improve this in the future)
	if _, errCreate := db.Create("wallets", wallet.ToMap()); errCreate != nil {
		if _, errUpdate := db.Change(wallet.Id, wallet.ToMap()); errUpdate != nil {
			panic(errUpdate)
		}
	}

	SaveTransactions(wallet.Transactions)
	return wallet
}

func SaveTransactions(transactions []investor.Transaction) []investor.Transaction {
	for _, transaction := range transactions {
		SaveTransaction(transaction)
	}

	return transactions
}

// func GetByClientEmail(email string) investor.Wallet {
// 	db := utils.GetDatabase()

// 	return wallet
// }
