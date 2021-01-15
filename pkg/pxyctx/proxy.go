package pxyctx

func RandomPxy() string {

	for pxy, ok := range Pool {
		if ok {
			return pxy
		}
	}

	return ""
}
