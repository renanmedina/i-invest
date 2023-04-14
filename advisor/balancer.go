package advisor

import (
	"investment-warlock/investor"
)

type BalanceSetup struct {
	RealState   float64
	Stocks      float64
	FixedIncome float64
}

type BalanceSuggestion struct {
	AssetKind         string
	CurrentPercentage float64
	CurrentTotal      float64
	TargetPercentage  float64
	Operation         string
	OperationAmount   float64
}

func MakeBalanceSetup(realstate float64, stock float64, fixed_income float64) BalanceSetup {
	return BalanceSetup{realstate / 100, stock / 100, fixed_income / 100}
}

func NewBalanceSuggestion(kind string, currentPercent float64, currentTotal float64, targetPercentage float64, suggestionAmount float64) BalanceSuggestion {
	operation := "MANTER"
	if suggestionAmount < 0 {
		operation = "VENDER"
	} else if suggestionAmount > 0 {
		operation = "COMPRAR"
	}

	return BalanceSuggestion{kind, currentPercent, currentTotal, targetPercentage, operation, suggestionAmount}
}

func BalanceWalletByAssetType(wallet investor.Wallet, setup BalanceSetup) []BalanceSuggestion {
	var suggestions []BalanceSuggestion
	walletTotal := wallet.Total()

	if setup.Stocks > 0 {
		kindTotalAmount := wallet.TotalForAssetKind("stock")
		targetAmount := walletTotal * setup.Stocks
		currentPercentage := 0.0
		if kindTotalAmount > 0 {
			currentPercentage = (kindTotalAmount / walletTotal) * 100
		}
		suggestionAmount := targetAmount - kindTotalAmount
		suggestion := NewBalanceSuggestion("Ações", currentPercentage, kindTotalAmount, setup.Stocks*100, suggestionAmount)
		suggestions = append(suggestions, suggestion)
	}

	if setup.RealState > 0 {
		kindTotalAmount := wallet.TotalForAssetKind("real_state")
		targetAmount := walletTotal * setup.RealState
		currentPercentage := 0.0
		if kindTotalAmount > 0 {
			currentPercentage = (kindTotalAmount / walletTotal) * 100
		}
		suggestionAmount := targetAmount - kindTotalAmount
		suggestion := NewBalanceSuggestion("Fundos Imobiliários", currentPercentage, kindTotalAmount, setup.RealState*100, suggestionAmount)
		suggestions = append(suggestions, suggestion)
	}

	if setup.FixedIncome > 0 {
		kindTotalAmount := wallet.TotalForAssetKind("fixed_income")
		targetAmount := walletTotal * setup.RealState
		currentPercentage := 0.0
		if kindTotalAmount > 0 {
			currentPercentage = (kindTotalAmount / walletTotal) * 100
		}
		suggestionAmount := targetAmount - kindTotalAmount
		suggestion := NewBalanceSuggestion("Renda Fixa", currentPercentage, kindTotalAmount, setup.FixedIncome*100, suggestionAmount)
		suggestions = append(suggestions, suggestion)
	}

	return suggestions
}
