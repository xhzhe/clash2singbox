package convert

import (
	"github.com/xmdhs/clash2singbox/clash"
	"github.com/xmdhs/clash2singbox/singbox"
)

func httpOpts(p *clash.Proxies, s *singbox.SingBoxOut) error {
	tls(p, s)
	p.Username = s.Username
	return nil
}

func socks5(p *clash.Proxies, s *singbox.SingBoxOut) error {
	tls(p, s)
	p.Username = s.Username
	if !p.Udp {
		s.Network = "tcp"
	}
	return nil
}
