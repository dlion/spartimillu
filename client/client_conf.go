package client

type SpartimilluClientConf struct {
	scheme string
	ip     string
	port   int
}

func NewSpartimilluClientConf(scheme, ip string, port int) SpartimilluClientConf {
	return SpartimilluClientConf{scheme: scheme, ip: ip, port: port}
}
