// 爬取网页并进行转码操作
package fetcher

import (
	"bufio"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"strings"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	var utf8Reader *transform.Reader
	if strings.Contains(url,"http://album.zhenai.com/u/") {
		client := &http.Client{}

		// 提交请求
		reqest, err := http.NewRequest("GET",url,nil)

		// 增加header选项
		reqest.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
		//reqest.Header.Add("Accept-Encoding","gzip, deflate")
		reqest.Header.Add("Connection","keep-alive")
		reqest.Header.Add("Accept-Language","zh-CN,zh;q=0.9")
		reqest.Header.Add("Cookie","sid=9c83c6ef-6d1d-4ea0-852f-a2dfd193ed9b; ipCityCode=-1; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1547283683,1548656253,1548724938; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1548734415")
		reqest.Header.Add("Host","album.zhenai.com")
		reqest.Header.Add("Referer",url)
		reqest.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.92 Safari/537.36")

		if err != nil {
			return nil,err
		}

		// 处理返回结果
		resp,_ := client.Do(reqest)
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

		utf8Reader = transform.NewReader(bodyReader,
			e.NewDecoder())

	} else {
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

		utf8Reader = transform.NewReader(bodyReader,
			e.NewDecoder())

	}
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
