package wallets

import (
	"github.com/renanmedina/i-invest/utils"
)

type WalletRepository struct {
	db *utils.DatabaseAdapdater
}

func NewWalletRepository() *WalletRepository {
	return &WalletRepository{
		db: utils.GetDatabase(),
	}
}

func (r *WalletRepository) GetById(walletId string) (*Wallet, error) {
	return &Wallet{}, nil
}

func (r *WalletRepository) Save(wallet Wallet) Wallet {
	// db := utils.GetDatabase()

	// try creating if fails probably exists then updated (should improve this in the future)
	// if _, errCreate := db.Create("wallets", wallet.ToMap()); errCreate != nil {
	// 	if _, errUpdate := db.Change(wallet.Id, wallet.ToMap()); errUpdate != nil {
	// 		panic(errUpdate)
	// 	}
	// }

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
