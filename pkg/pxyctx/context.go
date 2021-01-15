package pxyctx

var (
	PxyCandidateCh = make(chan string, 10)
	PxyReadyCh     = make(chan string, 10)
	PxyExpiredCh   = make(chan string, 10)
)

type Proxy struct {
	Address  string
	Port     string
	Protocol string
}

var Pool = make(map[string]bool)
