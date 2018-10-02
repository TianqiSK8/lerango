package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

//获得页面编码
func determineEncoding(reader *bufio.Reader) encoding.Encoding {
	bytes, err := reader.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}


//var rateLimiter = time.Tick(10 * time.Millisecond)
func Fetch(url string) ([]byte, error) {
	//<- rateLimiter
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	//resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	newReader := bufio.NewReader(resp.Body)
	e := determineEncoding(newReader)
	utf8reader := transform.NewReader(newReader, e.NewDecoder())
	//fmt.Printf("Fetching url: %s\n", url)
	return ioutil.ReadAll(utf8reader)
}
