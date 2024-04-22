package status_invest

import (
	"strconv"

	"github.com/renanmedina/investment-warlock/internal/market"
)

const (
	RELEVANT_FACT_TYPE = "Fato Relevante"
)

type AnnouncementService struct {
	apiClient market.ApiClient[AnnouncementApiResults]
}

type AnnouncementApiResults struct {
	Results []AnnouncementApiItem `json:"data"`
}

type AnnouncementApiItem struct {
	ReferenceDate    string `json:"dataReferencia"`
	Year             int    `json:"year"`
	Rank             int    `json:"rank"`
	AnnouncementType string `json:"especie"`
	Subject          string `json:"assunto"`
	LinkUrl          string `json:"linkPdf"`
}

func NewAnnouncementsService() *AnnouncementService {
	client := market.NewApiClient[AnnouncementApiResults](market.ApiConfig{
		ApiUrl: "https://statusinvest.com.br",
	})

	return &AnnouncementService{client}
}

func (s *AnnouncementService) GetByTickerCodeAndYear(tickerCode string, year int) (AnnouncementApiResults, error) {
	params := map[string]string{
		"year": strconv.Itoa(year),
		"code": tickerCode,
	}

	results, err := s.apiClient.Get("/acao/getassetreports", params)

	if err != nil {
		return AnnouncementApiResults{}, err
	}

	return *results, nil
}
