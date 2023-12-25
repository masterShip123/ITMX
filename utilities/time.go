package utilities

import (
	"time"
)

func GetTimeNowGMT() time.Time {
	return time.Now().UTC().Add(7 * time.Hour)
	//return GetTimeNowBangkok()
}
