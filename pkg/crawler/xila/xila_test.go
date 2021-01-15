package xiladaili

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_craw(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	craw(2)
}

func Test_pTrim(t *testing.T) {
	s := "https代理"
	s = pTrim(s)
	fmt.Println(s)
}
