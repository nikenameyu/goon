package public

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpPar struct {
	Url 	string
	Timeout int
	Follow  bool
	Body    string
	Header  [][2]string
}

func NewHttpPar() *HttpPar{
	return &HttpPar{
		"",
		10,
		true,
		"",
		[][2]string{},
	}
}

func HttpDoGet2Body(par *HttpPar) (http.Header, []byte, int){
	/* 跳过https验证 */
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout: time.Duration(par.Timeout) * time.Second,
	}
	if !par.Follow{
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return fmt.Errorf("first response")
		}
	}
	req, err := http.NewRequest("GET", par.Url, nil)
	if err != nil {
		return nil,nil,0
	}
	req.Header.Set("User-agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")

	if len(par.Header)!=0{
		for _,h := range(par.Header){
			hk, hv := h[0], h[1]
			req.Header.Set(hk, hv)
		}
	}

	resp, err := client.Do(req)
	if err!=nil{
		return nil,nil,0
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil,nil,0
	}
	return resp.Header,body,resp.StatusCode
}

//func HttpDoGet2Body(url string,timeout int,cookie string,follow bool) (http.Header, []byte, int){
//	/* 跳过https验证 */
//	tr := &http.Transport{
//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//	}
//	client := &http.Client{
//		Transport: tr,
//		Timeout: time.Duration(timeout) * time.Second,
//	}
//	if !follow{
//		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
//			return fmt.Errorf("first response")
//		}
//	}
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return nil,nil,0
//	}
//	req.Header.Set("User-agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
//	req.Header.Set("Accept", "*/*")
//	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
//	req.Header.Set("Cookie", cookie)
//	resp, err := client.Do(req)
//	if err!=nil{
//		return nil,nil,0
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil{
//		return nil,nil,0
//	}
//	return resp.Header,body,resp.StatusCode
//}

// 返回http请求code
func HttpDoGet2Code(url string,timeout int) int{
	/* 跳过https验证 */
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{
		Transport: tr,
		Timeout: time.Duration(timeout) * time.Second,
	}
	resp, err := c.Get(url)
	if err != nil {
		return 50000
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

// 返回http请求是否存活
func HttpDoGet2Alive(url string,timeout int) bool{
	/* 跳过https验证 */
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{
		Transport: tr,
		Timeout: time.Duration(timeout) * time.Second,
	}
	_, err := c.Get(url)
	if err != nil {
		return false
	}
	return true
}

func HttpDoHead2Resp(url string, timeout int) *http.Response{
	/* 跳过https验证 */
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{
		Transport: tr,
		Timeout: time.Duration(timeout) * time.Second,
	}
	resp, err := c.Head(url)
	if err!=nil{
		return nil
	}
	return resp
}

func HttpDoPost2Body(url string,timeout int,cookie string,follow bool,h http.Header, b string) (http.Header, []byte, int){
	/* 跳过https验证 */
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout: time.Duration(timeout) * time.Second,
	}
	if !follow{
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return fmt.Errorf("first response")
		}
	}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil,nil,0
	}
	req.Header.Set("User-agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1468.0 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)
	if err!=nil{
		return nil,nil,0
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil,nil,0
	}
	return resp.Header,body,resp.StatusCode
}