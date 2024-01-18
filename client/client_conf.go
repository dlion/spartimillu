package client

type SpartimilluClientConf struct {
	addresses []string
}

func NewSpartimilluClientConf(addresses []string) SpartimilluClientConf {
	return SpartimilluClientConf{addresses: addresses}
}
