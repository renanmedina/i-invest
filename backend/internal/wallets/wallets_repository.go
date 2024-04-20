package wallets

import (
	"github.com/renanmedina/investment-warlock/backend/utils"
	"github.com/surrealdb/surrealdb.go"
)

type WalletRepository struct {
	db *surrealdb.DB
}

func NewWalletRepository() *WalletRepository {
	return &WalletRepository{
		db: utils.GetDatabase(),
	}
}

func (r *WalletRepository) Save(wallet Wallet) Wallet {
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

func SaveTransactions(transactions []Transaction) []Transaction {
	for _, transaction := range transactions {
		SaveTransaction(transaction)
	}

	return transactions
}

// func GetByClientEmail(email string) Wallet {
// 	db := utils.GetDatabase()

// 	return wallet
// }
