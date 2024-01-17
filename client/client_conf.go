package client

type SpartimilluClientConf struct {
	scheme string
	ip     string
}

func NewSpartimilluClientConf(scheme, ip string) SpartimilluClientConf {
	return SpartimilluClientConf{scheme: scheme, ip: ip}
}
