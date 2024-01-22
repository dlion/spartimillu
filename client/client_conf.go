package client

type SpartimilluClientConf struct {
	addresses           []string
	healthcheckEndpoint string
}

func NewSpartimilluClientConf(addresses []string, healthcheckEndpoint string) SpartimilluClientConf {
	return SpartimilluClientConf{addresses: addresses, healthcheckEndpoint: healthcheckEndpoint}
}
