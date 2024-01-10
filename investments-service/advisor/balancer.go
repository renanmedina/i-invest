package advisor

import (
	"github.com/renanmedina/investment-warlock/investments-service/investor"
)

type BalanceSetup struct {
	Kind       string
	Percentage float64
}

type BalanceSuggestion struct {
	AssetKind         string
	CurrentPercentage float64
	CurrentTotal      float64
	TargetPercentage  float64
	Operation         string
	OperationAmount   float64
}

func MakeBalanceSetup(realstate float64, stock float64, fixed_income float64) []BalanceSetup {
	return []BalanceSetup{
		BalanceSetup{Kind: "real_state", Percentage: realstate / 100},
		BalanceSetup{Kind: "stock", Percentage: stock / 100},
		BalanceSetup{Kind: "fixed_income", Percentage: fixed_income / 100},
	}
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

func BalanceWalletByAssetType(wallet investor.Wallet, setups []BalanceSetup) []BalanceSuggestion {
	var suggestions []BalanceSuggestion
	walletTotal := wallet.Total()

	for _, setup := range setups {
		kindTotalAmount := wallet.TotalForAssetKind(setup.Kind)
		targetAmount := walletTotal * setup.Percentage
		currentPercentage := 0.0
		if kindTotalAmount > 0 {
			currentPercentage = (kindTotalAmount / walletTotal) * 100
		}

		suggestionAmount := targetAmount - kindTotalAmount
		suggestion := NewBalanceSuggestion(setup.Kind, currentPercentage, kindTotalAmount, setup.Percentage*100, suggestionAmount)
		suggestions = append(suggestions, suggestion)
	}

	return suggestions
}
