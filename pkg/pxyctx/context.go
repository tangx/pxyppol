package pxyctx

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
