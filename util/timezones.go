package util

import (
	"fmt"
	"time"
)

func ConvertTimezone(timeStr, srcZoneStr, destZoneStr string) (time.Time, error) {
	srcZone, err := time.LoadLocation(srcZoneStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse src zone, please use IANA timezone names %v", err)
	}
	destZone, err := time.LoadLocation(destZoneStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse dest zone, please use IANA timezone names %v", err)
	}
	var t time.Time
	if timeStr == "now" {
		t = time.Now()
	} else {
		t, err := time.Parse(time.Kitchen, timeStr)
		if err != nil {
			return t, err
		}
	}
	t = t.In(srcZone)
	return t.In(destZone), nil
}
