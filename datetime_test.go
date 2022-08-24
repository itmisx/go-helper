package helper

import (
	"fmt"
	"testing"
)

func TestToDate(*testing.T) {
	date := (&Datetime{}).WithTimezone(9).ToDate("Y-m-d H-i-s")
	fmt.Println(date)
}

func TestToTime(*testing.T) {
	tm, err := (&Datetime{}).WithTimezone(8).ToTime("2022-08-24 10:54:00")
	fmt.Println(tm, err)
}
