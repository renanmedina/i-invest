package wallets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/renanmedina/investment-warlock/utils"
)

type Wallet struct {
	Id            string        `json:"id,omitempty"`
	Name          string        `json:"name"`
	Client        Client        `json:"client"`
	Transactions  []Transaction `json:"transactions"`
	Consolidation map[string]ConsolidatorItem
}

func NewWallet(id string, name string, clientName string, transactions []Transaction) Wallet {
	return Wallet{
		Id:           id,
		Name:         name,
		Client:       NewClient(id, clientName, "renan@silvamedina.com.br"),
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

func (w Wallet) Quantity() int {
	quantity := 0

	for _, transaction := range w.Transactions {
		quantity += transaction.Quantity
	}

	return quantity
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

func (w Wallet) TransactionsMap() []map[string]interface{} {
	transactions := []map[string]interface{}{}

	for _, transaction := range w.Transactions {
		transactions = append(transactions, transaction.ToMap())
	}

	fmt.Println(transactions)
	return transactions
}

func (w *Wallet) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":             fmt.Sprintf("wallets:%s", w.Client.Email),
		"current_total":  w.Total(),
		"total_invested": w.TotalInvested(),
		"quantity":       w.Quantity(),
		"client":         w.Client.ToMap(),
	}
}

func (wallet *Wallet) PrintWalletHeader() {
	fmt.Println("===========================================================")
	fmt.Println("Wallet: ", wallet.Name)
	fmt.Println("Cliente: ", wallet.Client.Name)
	fmt.Println("Patrimonio atual: ", utils.CurrencyFormat(wallet.Total()))
	fmt.Println("Investimento realizado: ", utils.CurrencyFormat(wallet.TotalInvested()))
	fmt.Println("% variação: ", utils.PercentageFormat(wallet.VariationPercentage()))
}
