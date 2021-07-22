package Handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Proxy struct {
	ip   string
	port string
}

// GetProxiesFromUrl returns an arraylist with type Proxy
// Provided API needs to have Proxy:ip format per line
func GetProxiesFromUrl(siteUrl string) []Proxy {
	resp, err := http.Get(siteUrl)
	if err != nil {
		fmt.Println(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	proxies := strings.Split(string(bytes), "\n")
	return LoadProxies(proxies)
}

func LoadProxies(proxies []string) []Proxy {
	var proxyListArray []Proxy
	//proxies = proxies[:len(proxies) - 1]
	for _, element := range proxies {
		singleProxy := strings.Split(element, ":")
		if element != "" {
			proxyListArray = append(proxyListArray, Proxy{
				ip:   singleProxy[0],
				port: strings.Trim(singleProxy[1], "\r"),
			})
		}
	}
	return proxyListArray
}
func LoadNextProxy(proxyList []Proxy, index int) Proxy {
	return Proxy{
		ip:   proxyList[index].ip,
		port: proxyList[index].port,
	}
}
