package notification

import (
	"fmt"

	"github.com/renanmedina/investment-warlock/utils"
	twilio "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SmsService struct {
	smsClient      *twilio.RestClient
	smsAccountSSID string
	logger         *utils.ApplicationLogger
}

func NewSmsService() *SmsService {
	appConfigs := utils.GetConfigs()
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		AccountSid: appConfigs.TWILIO_ACCOUNT_SSID,
		Username:   appConfigs.TWILIO_API_SID,
		Password:   appConfigs.TWILIO_API_SECRET,
	})

	return &SmsService{
		client,
		appConfigs.TWILIO_SMS_SERVICE_SSID,
		utils.GetApplicationLogger(),
	}
}

func (s *SmsService) Send(toPhoneNumber string, msg string) error {
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(toPhoneNumber)
	params.SetBody(msg)
	params.MessagingServiceSid = &s.smsAccountSSID

	s.logger.Info(fmt.Sprintf("Sending SMS with twilio service to %s", toPhoneNumber))

	_, err := s.smsClient.Api.CreateMessage(params)

	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed sending SMS with twilio service to %s", toPhoneNumber))
		return err
	}

	s.logger.Info(fmt.Sprintf("Successfully sent SMS with twilio service to %s", toPhoneNumber))
	return nil
}
