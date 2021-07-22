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

// StartDownload /*func StartDownload(siteUrl string, Proxy Proxy){
func StartDownload(siteUrl string, proxy Proxy) {
	prepareDownload(siteUrl)
	partialChunk := rangeSize / 12
	chunkStart := 0
	chunkEnd := partialChunk
	outFile := make(chan string)
	var msg string
	for i := 0; i < 12; i++ {
		go DownloadFile(siteUrl, chunkStart, chunkEnd, proxy, outFile)
		chunkStart += partialChunk + 1
		chunkEnd += partialChunk
		msg += <-outFile
	}
	fileName := strings.Split(siteUrl, ":")[0]
	transformToFile(msg, fileName)
}

// DownloadFile /*func DownloadFile(siteUrl string, chunkStart int, chunkEnd int, proxy Proxy, out chan string){
func DownloadFile(siteUrl string, chunkStart int, chunkEnd int, proxy Proxy, out chan string) {
	defer close(out)
	request, _ := http.NewRequest("GET", siteUrl, nil)
	request.Header.Set("Range", "bytes="+strconv.Itoa(chunkStart)+"-"+strconv.Itoa(chunkEnd))
	proxyUrl, _ := url.Parse("https://" + proxy.port + ":" + proxy.port)
	transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	client := &http.Client{Transport: transport}
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
