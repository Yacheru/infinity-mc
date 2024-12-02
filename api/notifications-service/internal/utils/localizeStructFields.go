package utils

import (
	"fmt"
	"notifications-service/internal/entities"
	"notifications-service/pkg/constants"
	"strconv"
)

func LocalizeStruct(s *entities.Message) *entities.Message {
	durationInt, _ := strconv.Atoi(s.Duration)

	lastDigit := durationInt % 10
	lastTwoDigits := durationInt % 100

	switch {
	case lastDigit == 1 && lastTwoDigits != 11:
		s.Duration = fmt.Sprintf("%d месяц", durationInt)
	case lastDigit >= 2 && lastDigit <= 4 && (lastTwoDigits < 12 || lastTwoDigits > 14):
		s.Duration = fmt.Sprintf("%d месяца", durationInt)
	default:
		s.Duration = fmt.Sprintf("%d месяцев", durationInt)
	}

	switch s.Service {
	case constants.Nickname:
		s.Service = "никнейм"
	case constants.Hronon:
		s.Service = "хронон"
	case constants.Badge:
		s.Service = "значок"
	}

	return s
}
