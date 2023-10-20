package api_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/castus/lights-api/api"
)

func TestAfterNoon(t *testing.T) {
	ti := GetUnixTimestampForTime(21, 00)
	if api.Get(ti) != "250" {
		t.Fail()
	}
}

func TestBeforeNoon(t *testing.T) {
	ti := GetUnixTimestampForTime(11, 00)
	if api.Get(ti) != "917" {
		t.Fail()
	}
}

func TestNoon(t *testing.T) {
	ti := GetUnixTimestampForTime(12, 00)
	if api.Get(ti) != "1000" {
		t.Fail()
	}
}

func TestMidnight(t *testing.T) {
	ti := GetUnixTimestampForTime(00, 00)
	fmt.Println(api.Get(ti))
	if api.Get(ti) != "0" {
		t.Fail()
	}
}

func GetUnixTimestampForTime(hour int, minute int) time.Time {
	return time.Date(2023, time.October, 19, hour, minute, 0, 0, time.UTC)
}
