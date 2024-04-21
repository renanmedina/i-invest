package b3

import (
	"strconv"
	"strings"

	"github.com/thedatashed/xlsxreader"
)

const (
	STOCKS_SHEET_INDEX     = 0
	BDR_SHEET_INDEX        = 1
	REAL_STATE_SHEET_INDEX = 2
	TREASURE_SHEET_INDEX   = 3
)

type B3SummaryReportItem struct {
	ProductName      string
	Type             string
	BrokerName       string
	AccountNumber    string
	TickerCode       string
	CompanyId        string
	BankHolder       string
	Quantity         float64
	ClosingUnitPrice float64
	TotalAmount      float64
}

func ParseB3SummaryReport(filepath string) ([]B3SummaryReportItem, error) {
	xl, err := xlsxreader.OpenFile(filepath)
	defer xl.Close()

	if err != nil {
		return nil, err
	}

	items := make([]B3SummaryReportItem, 0)
	stocksItems := parseAndBuildListFromSheet(*xl, STOCKS_SHEET_INDEX, "stock")
	bdrItems := parseAndBuildListFromSheet(*xl, BDR_SHEET_INDEX, "bdr")
	realStateItems := parseAndBuildListFromSheet(*xl, REAL_STATE_SHEET_INDEX, "real_state")
	treasuresItems := parseAndBuildListFromSheet(*xl, TREASURE_SHEET_INDEX, "treasure")

	items = append(items, stocksItems...)
	items = append(items, bdrItems...)
	items = append(items, realStateItems...)
	items = append(items, treasuresItems...)

	// byAssetTypeDictionary := map[string][]B3SummaryReportItem{
	// 	"stocks":     stocksItems,
	// 	"real_state": realStateItems,
	// 	"bdr":        bdrItems,
	// 	"treasure":   treasuresItems,
	// 	"all":        items,
	// }

	return items, nil
}

func parseAndBuildListFromSheet(xl xlsxreader.XlsxFileCloser, index int, assetType string) []B3SummaryReportItem {
	assetItems := make([]B3SummaryReportItem, 0)
	headerSkipped := false

	for row := range xl.ReadRows(xl.Sheets[index]) {
		if !headerSkipped || len(row.Cells) < 5 {
			headerSkipped = true
			continue
		}

		// check if ticker code is present to ignore invalid lines
		if row.Cells[0].Value == "" {
			continue
		}

		parsedItem := newB3SummaryFromXlsRow(row, assetType)
		assetItems = append(assetItems, parsedItem)
	}

	return assetItems
}

func newB3SummaryFromXlsRow(row xlsxreader.Row, assetType string) B3SummaryReportItem {
	switch assetType {
	case "bdr":
		return parseBDRFromXlsRow(row)
	case "treasure":
		return parseTreasureFromXlsRow(row)
	default:
		return parseDefaultAssetFromXlsRow(row, assetType)
	}
}

func parseBDRFromXlsRow(row xlsxreader.Row) B3SummaryReportItem {
	qtd, _ := strconv.ParseFloat(row.Cells[7].Value, 64)
	closing_price, _ := strconv.ParseFloat(row.Cells[11].Value, 64)
	final_price, _ := strconv.ParseFloat(row.Cells[12].Value, 64)

	return B3SummaryReportItem{
		ProductName:      strings.TrimSpace(row.Cells[0].Value),
		Type:             "bdr",
		BrokerName:       strings.TrimSpace(row.Cells[1].Value),
		AccountNumber:    strings.TrimSpace(row.Cells[2].Value),
		TickerCode:       strings.TrimSpace(row.Cells[3].Value),
		CompanyId:        "",
		BankHolder:       strings.TrimSpace(row.Cells[6].Value),
		Quantity:         qtd,
		ClosingUnitPrice: closing_price,
		TotalAmount:      final_price,
	}
}

func parseDefaultAssetFromXlsRow(row xlsxreader.Row, assetType string) B3SummaryReportItem {
	qtd, _ := strconv.ParseFloat(row.Cells[8].Value, 64)
	closing_price, _ := strconv.ParseFloat(row.Cells[12].Value, 64)
	final_price, _ := strconv.ParseFloat(row.Cells[13].Value, 64)

	return B3SummaryReportItem{
		ProductName:      strings.TrimSpace(row.Cells[0].Value),
		Type:             assetType,
		BrokerName:       strings.TrimSpace(row.Cells[1].Value),
		AccountNumber:    strings.TrimSpace(row.Cells[2].Value),
		TickerCode:       strings.TrimSpace(row.Cells[3].Value),
		CompanyId:        strings.TrimSpace(row.Cells[4].Value),
		BankHolder:       strings.TrimSpace(row.Cells[7].Value),
		Quantity:         qtd,
		ClosingUnitPrice: closing_price,
		TotalAmount:      final_price,
	}
}

func parseTreasureFromXlsRow(row xlsxreader.Row) B3SummaryReportItem {
	qtd, _ := strconv.ParseFloat(row.Cells[5].Value, 64)
	closing_price, _ := strconv.ParseFloat(row.Cells[9].Value, 64)
	final_price, _ := strconv.ParseFloat(row.Cells[12].Value, 64)

	return B3SummaryReportItem{
		ProductName:      strings.TrimSpace(row.Cells[0].Value),
		Type:             "bdr",
		BrokerName:       strings.TrimSpace(row.Cells[1].Value),
		AccountNumber:    strings.TrimSpace(row.Cells[2].Value),
		TickerCode:       strings.TrimSpace(row.Cells[3].Value),
		CompanyId:        "",
		BankHolder:       strings.TrimSpace(row.Cells[6].Value),
		Quantity:         qtd,
		ClosingUnitPrice: closing_price,
		TotalAmount:      final_price,
	}
}
