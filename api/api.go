package api

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func Get(unixTimestamp time.Time) string {
	now := time.Now()
	var secondsDay float64 = 86400
	startOfDay := unixTimestamp.Truncate(24 * time.Hour)
	noon := startOfDay.Add(12 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	fmt.Println("Start of day:", startOfDay)
	fmt.Println("End of day:", endOfDay)
	fmt.Println(now, secondsDay)

	var value float64
	if unixTimestamp.Before(noon) {
		subTime := unixTimestamp.Sub(startOfDay).Seconds()
		value = math.Ceil(subTime * 1000 / (secondsDay / 2))
	} else {
		subTime := endOfDay.Sub(unixTimestamp).Seconds()
		secs := (secondsDay / 2) - subTime
		value = 1000 - math.Ceil(secs*1000/(secondsDay/2))
	}

	return strconv.FormatFloat(value, 'f', 0, 64)
}
