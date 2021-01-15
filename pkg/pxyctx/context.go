package pxyctx

import (
	"fmt"
	"time"
)

var (
	PxyCandidateCh = make(chan string, 100)
	PxyReadyCh     = make(chan string, 100)
	PxyExpiredCh   = make(chan string, 100)
)

type Proxy struct {
	Address  string
	Port     string
	Protocol string
}

var Pool = make(map[string]bool)

func init() {
	Pool["http://127.0.0.1:7890"] = true
	Pool["socks5://127.0.0.1:7890"] = true
}

func init() {
	// go counter()
}

func counter() {

	for {
		time.Sleep(1 * time.Second)
		fmt.Println("PxyCandidateCh", len(PxyCandidateCh))
		fmt.Println("PxyReadyCh", len(PxyReadyCh))
		fmt.Println("PxyExpiredCh", len(PxyExpiredCh))
	}
}
