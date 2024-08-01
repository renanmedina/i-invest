package market

import (
	"strconv"

	"github.com/renanmedina/i-invest/internal/integration"
)

const (
	RELEVANT_FACT_TYPE = "Fato Relevante"
)

type AnnouncementService struct {
	apiClient integration.ApiClient[AnnouncementApiResults]
}

type AnnouncementApiResults struct {
	Items []AnnouncementApiItem `json:"data"`
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
	client := integration.NewApiClient[AnnouncementApiResults](integration.ApiConfig{
		ApiUrl: "https://statusinvest.com.br",
	})

	return &AnnouncementService{client}
}

func (s *AnnouncementService) GetByTickerCodeAndYear(tickerCode string, year int) (AnnouncementApiResults, error) {
	params := map[string]string{
		"year": strconv.Itoa(year),
		"code": tickerCode,
	}

	results, err := s.apiClient.Get("/acao/getassetreports", params, defaultHeaders())

	if err != nil {
		return AnnouncementApiResults{}, err
	}

	return *results, nil
}

func defaultHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   "PostmanRuntime/7.37.3",
		"Accept":       "*/*",
	}
}
