package helper

import (
	"strconv"
	"strings"
)

// CompareVersion 比较版本大小
func CompareVersion(prefix, version1, version2 string) int {
	version1 = strings.TrimPrefix(version1, prefix)
	version2 = strings.TrimPrefix(version2, prefix)
	if version1 == version2 {
		return 1
	}
	verArr1 := strings.Split(version1, ".")
	verArr2 := strings.Split(version2, ".")
	for verIndex, verSep := range verArr1 {
		if verIndex > len(verArr2)-1 {
			return 2
		}
		verSepInt1, err1 := strconv.Atoi(verSep)
		verSepInt2, err2 := strconv.Atoi(verArr2[verIndex])
		if err1 != nil || err2 != nil {
			return -1
		}
		if verSepInt1 < verSepInt2 {
			return 0
		} else if verSepInt1 == verSepInt2 {
			continue
		} else {
			return 2
		}
	}
	return 1
}
