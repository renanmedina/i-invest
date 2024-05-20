package announcements

import (
	"fmt"

	"github.com/renanmedina/investment-warlock/internal/accounts"
	"github.com/renanmedina/investment-warlock/internal/event_store"
	"github.com/renanmedina/investment-warlock/internal/feature_flags"
	"github.com/renanmedina/investment-warlock/internal/notification"
	"github.com/renanmedina/investment-warlock/utils"
)

type NotifyCompanyNewAnnouncementToUserHandler struct {
	featureFlag *feature_flags.FeatureFlag
}

func NewNotifyCompanyNewAnnouncementToUserHandler() *NotifyCompanyNewAnnouncementToUserHandler {
	allFeatureFlags := feature_flags.NewFeatureFlagsRepository()
	flag, err := allFeatureFlags.GetByFlagName(feature_flags.FLAG_NOTIFY_ANNOUNCEMENT_THROUGH_SMS)

	if err != nil {
		panic(err)
	}

	return &NotifyCompanyNewAnnouncementToUserHandler{flag}
}

func (h *NotifyCompanyNewAnnouncementToUserHandler) Handle(event event_store.PublishableEvent) {
	if !h.featureFlag.Enabled {
		return
	}

	use_case := NewNotifyCompanyNewAnnouncementsToUser()
	use_case.Execute(event.ObjectId())
}

type NotifyCompanyNewAnnouncementToUser struct {
	smsService       *notification.SmsService
	allAnnouncements *AnnouncementsRepository
	allUsers         *accounts.UsersRepository
	logger           *utils.ApplicationLogger
}

func NewNotifyCompanyNewAnnouncementsToUser() NotifyCompanyNewAnnouncementToUser {
	return NotifyCompanyNewAnnouncementToUser{
		notification.NewSmsService(),
		NewAnnouncementsRepository(),
		accounts.NewUsersRepository(),
		utils.GetApplicationLogger(),
	}
}

func (use_case *NotifyCompanyNewAnnouncementToUser) Execute(announcementId string) error {
	announcement, err := use_case.allAnnouncements.GetById(announcementId)

	// ignore notification for announcements that are not for today and prevent charging when filling up announcements
	if !announcement.ForToday() {
		return nil
	}

	if err != nil {
		use_case.logger.Error(err.Error())
		return err
	}

	user, err := use_case.allUsers.GetWithAnnoucementMonitoringEnabledByTickerCode(announcement.TickerCode)

	if err != nil {
		use_case.logger.Error(err.Error())
		return err
	}

	msg := formatAnnouncementMessage(&user, announcement)
	errSend := use_case.smsService.Send(user.PhoneNumber, msg)

	if errSend != nil {
		fmt.Println(errSend)
		use_case.logger.Error(errSend.Error())
		return errSend
	}

	return nil
}

func formatAnnouncementMessage(user *accounts.User, announcement *CompanyAnnouncement) string {
	smsBody := fmt.Sprintf("*Novo anúncio realizado pela empresa do ativo %s \r\n\r\n", announcement.TickerCode)
	smsBody += fmt.Sprintf("Título: %s \r\n\r\n", announcement.Subject)
	smsBody += fmt.Sprintf("Tipo: %s \r\n\r\n", announcement.AnnouncementType)
	smsBody += fmt.Sprintf("Url: %s \r\n\r\n", announcement.OriginalFileUrl)
	return smsBody
}
