package Model

import "ThunderHades/Handler"

type ProxyPool struct {
	allProxies       []Handler.Proxy
	availableProxies []Handler.Proxy
	badProxies       []Handler.Proxy
	bannedProxies    []Handler.Proxy
	proxyLocked      []bool
}

//region getter/setter functions
func (p *ProxyPool) GetAllProxies() []Handler.Proxy {
	return p.allProxies
}

func (p *ProxyPool) SetAllProxies(pool []Handler.Proxy) {
	p.allProxies = pool
}

func (p *ProxyPool) GetAvailableProxies() []Handler.Proxy {
	return p.availableProxies
}

func (p *ProxyPool) SetAvailableProxies(pool []Handler.Proxy) {
	p.availableProxies = pool
}

func (p *ProxyPool) GetBadProxies() []Handler.Proxy {
	return p.badProxies
}

func (p *ProxyPool) SetBadProxies(pool []Handler.Proxy) {
	p.badProxies = pool
}

func (p *ProxyPool) GetBannedProxies() []Handler.Proxy {
	return p.bannedProxies
}

func (p *ProxyPool) SetBannedProxies(pool []Handler.Proxy) {
	p.bannedProxies = pool
}

func (p *ProxyPool) GetLock() []bool {
	return p.proxyLocked
}

func (p *ProxyPool) SetLock(lock []bool) {
	p.proxyLocked = lock
}

//endregion
