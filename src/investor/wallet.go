package investor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Wallet struct {
	Id            int64         `json:"id"`
	Name          string        `json:"name"`
	Client        Client        `json:"client"`
	Transactions  []Transaction `json:"transactions"`
	Consolidation map[string]ConsolidatorItem
}

func NewWallet(id int64, name string, clientName string, transactions []Transaction) Wallet {
	return Wallet{
		Id:   id,
		Name: name,
		Client: Client{
			Id:   id,
			Name: clientName,
		},
		Transactions: transactions,
	}
}

func (w Wallet) Total() float64 {
	total := 0.0
	for _, consolidatorItem := range w.Consolidation {
		total += consolidatorItem.CurrentAmount
	}

	return total
}

func (w Wallet) TotalInvested() float64 {
	total := 0.0
	for _, transaction := range w.Transactions {
		total += transaction.Total()
	}

	return total
}

func (w Wallet) VariationPercentage() float64 {
	percentage := 0.0
	totalAmount := w.Total()
	totalInvested := w.TotalInvested()

	if totalAmount != 0 {
		differenceAmount := totalAmount - totalInvested
		fmt.Println(differenceAmount)
		percentage = (differenceAmount * 100) / totalInvested
	}

	return percentage
}

func (w Wallet) HasAsset(assetTicker string) bool {
	_, alreadyOnMap := w.Consolidation[assetTicker]
	return alreadyOnMap
}

func (w Wallet) GetConsolidation(assetTicker string) (ConsolidatorItem, bool) {
	consolidation, hasAsset := w.Consolidation[assetTicker]
	return consolidation, hasAsset
}

func (w Wallet) Consolidate() Wallet {
	w.Consolidation = ConsolidateByAsset(w)
	return w
}

func (w Wallet) TotalForAssetKind(targetKind string) float64 {
	total := 0.0
	for _, consolidatorItem := range w.Consolidation {
		if consolidatorItem.HasDetails(targetKind) {
			total += consolidatorItem.AverageAmount
		}
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

	wallet = wallet.Consolidate()
	return wallet
}
