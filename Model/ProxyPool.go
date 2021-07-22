package Model

import "ThunderHades/Handler"

type ProxyPool struct {
	availableProxies []Handler.Proxy
	aliveProxies     []Handler.Proxy
	deadProxies      []Handler.Proxy
	badProxies       []Handler.Proxy
	bannedProxies    []Handler.Proxy
	proxyLocked      bool
}

func (p ProxyPool) GetAvailableProxies() []Handler.Proxy {
	return p.availableProxies
}
func (p ProxyPool) SetAvailableProxies(pool []Handler.Proxy) {
	p.aliveProxies = pool
}

func (p ProxyPool) GetAliveProxies() []Handler.Proxy {
	return p.availableProxies
}
func (p ProxyPool) SetAliveProxies(pool []Handler.Proxy) {
	p.aliveProxies = pool
}

func (p ProxyPool) GetDeadProxies() []Handler.Proxy {
	return p.availableProxies
}
func (p ProxyPool) SetDeadProxies(pool []Handler.Proxy) {
	p.aliveProxies = pool
}

func (p ProxyPool) GetBadProxies() []Handler.Proxy {
	return p.availableProxies
}
func (p ProxyPool) SetBadProxies(pool []Handler.Proxy) {
	p.aliveProxies = pool
}

func (p ProxyPool) GetBannedProxies() []Handler.Proxy {
	return p.availableProxies
}
func (p ProxyPool) SetBannedProxies(pool []Handler.Proxy) {
	p.aliveProxies = pool
}

//region getter/setter proxy lock
func (p *ProxyPool) GetLock() bool {
	return p.proxyLocked
}
func (p *ProxyPool) SetLock(lock bool) {
	p.proxyLocked = lock
}

//endregion
