package B3

import (
	"github.com/renanmedina/i-invest/internal/market"
	"github.com/renanmedina/i-invest/utils"
)

const API_URL = "https://investidor.b3.com.br"

type WalletService struct {
	negotiationsApi  market.ApiClient[B3ApiResult[NegotiationDayItem]]
	consolidationApi market.ApiClient[B3ApiResult[ConsolidatedByProductItem]]
}

func NewWalletServiceWithToken() *WalletService {
	configs := utils.GetConfigs()
	return NewWalletService(configs.B3_API_TOKEN)
}

func NewWalletService(apiToken string) *WalletService {
	negotiationsApi := market.NewApiClient[B3ApiResult[NegotiationDayItem]](market.ApiConfig{ApiUrl: API_URL, AuthToken: apiToken, LogEnabled: true})
	consolidationApi := market.NewApiClient[B3ApiResult[ConsolidatedByProductItem]](market.ApiConfig{ApiUrl: API_URL, AuthToken: apiToken, LogEnabled: true})

	return &WalletService{
		negotiationsApi,
		consolidationApi,
	}
}

func (s *WalletService) GetNegotiationsByPeriod(dateStart string, dateEnd string) ([]NegotiationDayItem, error) {
	params := map[string]string{
		"dataInicio": dateStart,
		"dataFim":    dateEnd,
	}

	response, err := s.negotiationsApi.Get("/api/extrato-negociacao-ativos/v1/negociacao-ativos/1",
		params,
		make(map[string]string),
	)

	if err != nil {
		return nil, err
	}

	return response.Itens, nil
}

func (s WalletService) GetConsolidatedSnapshotByDate(date string) ([]ConsolidatedByProductItem, error) {
	params := map[string]string{"data": date}

	response, err := s.consolidationApi.Get("/api/extrato-posicao/v2/posicao",
		params,
		make(map[string]string),
	)

	if err != nil {
		return nil, err
	}

	return response.Itens, nil
}

type B3ApiResult[T any] struct {
	CurrentPage int `json:"paginaAtual"`
	TotalPages  int `json:"totalPaginas"`
	Itens       []T `json:"itens"`
}

type TransactionDayItem struct {
	Date         string            `json:"data"`
	Transactions []TransactionItem `json:"movimentacoes"`
}

type TransactionItem struct {
	Date            string  `json:"dataMovimentacao"`
	OperationType   string  `json:"tipoOperacao"`
	TransactionType string  `json:"tipoMovimentacao"`
	ProductName     string  `json:"nomeProduto"`
	Quantity        float32 `json:"quantidade"`
	Amount          float64 `json:"valorOperacao"`
	UnitPrice       float64 `json:"precoUnitario"`
}

type NegotiationDayItem struct {
	Date         string            `json:"data"`
	Negotiations []NegotiationItem `json:"negociacaoAtivos"`
}

type NegotiationItem struct {
	Date          string
	MarketName    string  `json:"mercado"`
	OperationType string  `json:"tipoMovimentacao"`
	TickerCode    string  `json:"codigoNegociacao"`
	Quantity      float32 `json:"quantidade"`
	Amount        float64 `json:"valor"`
	UnitPrice     float64 `json:"preco"`
}

func (item *NegotiationItem) AssetType() string {
	if item.MarketName == "Mercado Fracionário" {
		return "stock"
	}

	return item.MarketName
}

type ConsolidatedByProductItem struct {
	ProductType string                     `json:"tipoProduto"`
	Positions   []ConsolidatedPositionItem `json:"posicoes"`
}

type ConsolidatedPositionItem struct {
	Ticker      string  `json:"codigoNegociacao"`
	ProductName string  `json:"produto"`
	CompanyName string  `json:"razaoSocial"`
	Quantity    float32 `json:"quantidade"`
	TotalAmount float64 `json:"valorAtualizado"`
	UnitPrice   float64 `json:"precoFechamento"`
}
