package investor

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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
	for _, transaction := range w.Transactions {
		total += transaction.Total()
	}

	return total
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
			total += consolidatorItem.TotalCost
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

func ImportFromCsv(filepath string) (Wallet, error) {
	csvFile, err := os.Open(filepath)
	transactions := []Transaction{}

	if err != nil {
		return Wallet{}, err
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	transactionsData, err := csvReader.ReadAll()

	if err != nil {
		return Wallet{}, err
	}

	for rowIndex, line := range transactionsData {
		// ignore header
		if rowIndex > 0 {
			assetType := "stock"
			if line[2] == "Mercado à Vista" {
				assetType = "real_state"
			}

			quantity, _ := strconv.Atoi(line[6])
			replacedPrice := strings.ReplaceAll(strings.ReplaceAll(line[7], "R$", ""), " ", "")
			price, _ := strconv.ParseFloat(replacedPrice, 64)

			if line[1] == "Venda" {
				quantity *= -1
			}

			transaction := NewTransaction(
				assetType,
				line[5],
				price,
				quantity,
				0.0,
				line[0],
			)

			transactions = append(transactions, transaction)
		}
	}

	wallet := NewWallet(1, "Wallet de testes", "Renan Medina", transactions).Consolidate()
	return wallet, nil
}
