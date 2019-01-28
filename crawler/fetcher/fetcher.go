// 爬取网页并进行转码操作
package fetcher

import (
	"bufio"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("Wrong status code: %d ", resp.StatusCode)
	}

	// 转换获取页面的数据编码
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	// utf8Reader := transform.NewReader(resp.Body,
	// simplifiedchinese.GBK.NewDecoder())

	utf8Reader := transform.NewReader(resp.Body,
		e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// 判断页面编码，并返回encoding
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	// 因为resp.Body只能读取一次，因此创建一个buffer把数据拷贝一遍
	// 因为DetermineEncoding接收的content参数大小为1024
	bytes, err := r.Peek(1024)
	if err != nil {
		// 出错返回默认编码 utf8
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
