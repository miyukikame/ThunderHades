package Handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var file string
var rangeSize int

func prepareDownload(siteUrl string) {
	resp, _ := http.Get(siteUrl)
	rangeHeader := resp.Header.Get("Content-Length")
	rangeSize, _ = strconv.Atoi(rangeHeader)
}
func TestProxy(siteUrl string) {
	var proxyList = GetProxiesFromUrl("https://api.proxyscrape.com/v2/?request=getproxies&protocol=http&timeout=150&country=all&ssl=yes&anonymity=elite&simplified=true")
	proxyUrl, _ := url.Parse("https://" + proxyList[1].ip + ":" + proxyList[1].port)
	rawSiteUrl, _ := url.Parse(siteUrl)
	transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{Transport: transport}
	request, _ := http.NewRequest("GET", rawSiteUrl.String(), nil)
	response, _ := client.Do(request)
	htmldata, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(htmldata))
}

// StartDownload /*func StartDownload(siteUrl string, proxy proxy){
func StartDownload(siteUrl string) {
	prepareDownload(siteUrl)
	partialChunk := rangeSize / 12
	chunkStart := 0
	chunkEnd := partialChunk
	outFile := make(chan string)
	var msg string
	for i := 0; i < 12; i++ {
		go DownloadFile(siteUrl, chunkStart, chunkEnd, outFile)
		chunkStart += partialChunk + 1
		chunkEnd += partialChunk
		msg += <-outFile
	}
	fileName := strings.Split(siteUrl, ":")[0]
	transformToFile(msg, fileName)
}

// DownloadFile /*func DownloadFile(siteUrl string, chunkStart int, chunkEnd int, proxy proxy, out chan string){
func DownloadFile(siteUrl string, chunkStart int, chunkEnd int, out chan string) {
	defer close(out)
	request, _ := http.NewRequest("GET", siteUrl, nil)
	request.Header.Set("Range", "bytes="+strconv.Itoa(chunkStart)+"-"+strconv.Itoa(chunkEnd))
	client := &http.Client{}
	response, err := client.Do(request)
	if response.StatusCode > 350 {
		fmt.Println(err)
	}
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	out <- string(bodyBytes)
}
func transformToFile(in string, fileName string) {
	f, _ := os.Create("D:\\goProjects\\ThunderHades\\" + fileName)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	_, err := f.WriteString(in)
	if err != nil {
		return
	}
}
