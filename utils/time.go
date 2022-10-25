package utils

import (
	"sonarhook/config"
	"strings"
	"time"
)

func ParseTime(dateTime string) string {

	dateTimeN := strings.SplitN(dateTime, "+", -1)[0]

	loc, _ := time.LoadLocation(config.Timezone)

	time, _ := time.Parse("2006-01-02T15:04:05", dateTimeN)

	return time.In(loc).Format("2006-01-02 15:04:05")
}
