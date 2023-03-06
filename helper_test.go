package helper

import (
	"fmt"
	"testing"
)

func TestCompareVersion(t *testing.T) {
	fmt.Println(CompareVersion("v", "3.001.1", "3.1.2"))
	fmt.Println(CompareVersion("v", "3.1.1", "3.1"))
	fmt.Println(CompareVersion("v", "3.1.1", "3.1.1"))
	fmt.Println(CompareVersion("v", "3.1.2", "3.1.1"))
}
