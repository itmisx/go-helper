package helper

import (
	"fmt"
	"testing"
)

func TestDatatime(*testing.T) {
	fm := (&Datetime{}).WithTimezone(9).ToDate("Y-m-d H-i-s")
	fmt.Println(fm)
}
