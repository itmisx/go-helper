package helper

import (
	"fmt"
	"testing"
)

func TestHttps(*testing.T) {
	t, err := HttpsExpireTimestamp("http://www.baidu.com")
	fmt.Println(err, t)
}
