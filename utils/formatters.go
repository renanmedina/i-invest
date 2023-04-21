package utils

import (
	"fmt"
)

func currencyFormat(price float64) string {
	return fmt.Sprintf("R$ %v", numberFormat(price))
}

func percentageFormat(percent float64) string {
	return fmt.Sprintf("%v%%", numberFormat(percent))
}

func numberFormat(number float64) string {
	return fmt.Sprintf("%.2f", number)
}

func translateKind(kind string) string {
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
