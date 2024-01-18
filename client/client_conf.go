package client

type SpartimilluClientConf struct {
	address string
}

func NewSpartimilluClientConf(address string) SpartimilluClientConf {
	return SpartimilluClientConf{address: address}
}
