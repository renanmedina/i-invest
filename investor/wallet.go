package investor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Wallet struct {
	Id           int64         `json:"id"`
	Name         string        `json:"name"`
	Client       Client        `json:"client"`
	Transactions []Transaction `json:"transactions"`
}

func (w Wallet) Total() float64 {
	total := 0.0
	for _, transaction := range w.Transactions {
		total += transaction.Total()
	}

	return total
}

func BuildWalletFromJsonFile(filepath string) Wallet {
	jsonFile, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Erro ao ler o arquivo")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var wallet Wallet

	json.Unmarshal(byteValue, &wallet)

	return wallet
}
