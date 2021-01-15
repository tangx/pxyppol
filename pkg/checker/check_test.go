package checker

import (
	"fmt"
	"testing"
)

func Test_get(t *testing.T) {

	ok := get("http://ip.cip.cc", "http://112.111.77.101:9999")

	fmt.Println(ok)

}
