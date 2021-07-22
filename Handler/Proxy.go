package Handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type proxy struct {
	ip   string
	port string
}

// GetProxiesFromUrl returns an arraylist with type proxy
// Provided API needs to have proxy:ip format per line
func GetProxiesFromUrl(siteUrl string) []proxy {
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

func LoadProxies(proxies []string) []proxy {
	var proxyListArray []proxy
	//proxies = proxies[:len(proxies) - 1]
	for _, element := range proxies {
		singleProxy := strings.Split(element, ":")
		if element != "" {
			proxyListArray = append(proxyListArray, proxy{
				ip:   singleProxy[0],
				port: strings.Trim(singleProxy[1], "\r"),
			})
		}
	}
	return proxyListArray
}
func LoadNextProxy(proxyList []proxy, index int) proxy {
	return proxy{
		ip:   proxyList[index].ip,
		port: proxyList[index].port,
	}
}
