package myhttp

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string) (s string, ok bool){
	// TODO 传入目标url， 返回爬取的内容

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{ Transport : tr }

	req, err := http.NewRequest("GET", url, strings.NewReader("") )

	if err != nil {
		fmt.Println(err)
		return "", false
	}

	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:66.0) Gecko/20100101 Firefox/66.0")

	resp, err := client.Do(req) // 执行爬取

	if resp != nil { defer resp.Body.Close() }
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	body, err := ioutil.ReadAll(resp.Body) // 读取爬取内容

	if err != nil {
		fmt.Println(err)
		return "", false
	}

	return string(body), true
}

func Post(url string, postdata string) (s string, ok bool) {
	// TODO 传入目标url，和post数据，返回爬取的url

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{ Transport : tr }

	req, err := http.NewRequest("POST", url, strings.NewReader(postdata) )

	if err != nil {
		fmt.Println(err)
		return "", false
	}

	// 设置headers头
	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:66.0) Gecko/20100101 Firefox/66.0")
	req.Header.Set("Content-Type","application/x-www-form-urlencoded")

	resp, err := client.Do(req) // 执行爬取
	if resp != nil { defer resp.Body.Close() }
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return "", false
	}

	return string(body), true
}