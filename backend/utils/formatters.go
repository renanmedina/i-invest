package utils

import (
	"fmt"
)

func CurrencyFormat(price float64) string {
	return fmt.Sprintf("R$ %v", NumberFormat(price))
}

func PercentageFormat(percent float64) string {
	return fmt.Sprintf("%v%%", NumberFormat(percent))
}

func NumberFormat(number float64) string {
	return fmt.Sprintf("%.2f", number)
}

func TranslateKind(kind string) string {
	switch kind {
	case "fii":
		return "Fundos Imobiliários"
	case "real_state":
		return "Fundos Imobiliários"
	case "stock":
		return "Ações"
	case "bdr":
		return "BDR"
	}

	return kind
}
