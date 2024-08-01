package discord

import (
	"fmt"

	"github.com/renanmedina/i-invest/internal/integration"
	"github.com/renanmedina/i-invest/utils"
)

type MessageService struct {
	apiClient integration.ApiClient[MessageSentResult]
	channelId string
}

func NewMessageService() *MessageService {
	configs := utils.GetConfigs()

	return &MessageService{
		apiClient: integration.NewApiClient[MessageSentResult](integration.ApiConfig{
			ApiUrl:    "https://discord.com/api/v10",
			AuthToken: configs.DISCORD_BOT_TOKEN,
		}),
		channelId: configs.DISCORD_CHANNEL_ID,
	}
}

func (service *MessageService) Send(message string) error {
	return service.sendMessage(service.channelId, message)
}

func (service *MessageService) sendMessage(channelId string, message string) error {
	api := service.apiClient
	path := fmt.Sprintf("/channels/%s/messages", channelId)
	params := map[string]string{
		"content": message,
	}

	_, err := api.Post(path, params, map[string]string{})

	if err != nil {
		return err
	}

	return nil
}

type MessageSentResult struct{}
