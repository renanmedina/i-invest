package B3

import (
	"github.com/renanmedina/i-invest/internal/market"
)

const API_URL = "https://investidor.b3.com.br"

type WalletService struct {
	transactionsApi  market.ApiClient[B3ApiResult[TransactionDayItem]]
	consolidationApi market.ApiClient[B3ApiResult[ConsolidatedByProductItem]]
}

func NewWalletService(apiToken string) *WalletService {
	transactionsApi := market.NewApiClient[B3ApiResult[TransactionDayItem]](market.ApiConfig{ApiUrl: API_URL, AuthToken: apiToken, LogEnabled: true})
	consolidationApi := market.NewApiClient[B3ApiResult[ConsolidatedByProductItem]](market.ApiConfig{ApiUrl: API_URL, AuthToken: apiToken, LogEnabled: true})

	return &WalletService{
		transactionsApi,
		consolidationApi,
	}
}

func (s *WalletService) GetTransactionsByDate(dateStart string, dateEnd string) ([]TransactionDayItem, error) {
	params := map[string]string{
		"dataInicio": dateStart,
		"dataFim":    dateEnd,
	}

	response, err := s.transactionsApi.Get("/api/extrato-movimentacao/v2/movimentacao",
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
