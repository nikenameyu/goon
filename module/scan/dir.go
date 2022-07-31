package scan

import (
	"fmt"
	"goon3/public"
	"regexp"
	"runtime"
	"strings"
)

func DirScan(urls []string){
	runtime.GOMAXPROCS(runtime.NumCPU())
	input := make(chan string, len(urls))
	result := make(chan string, len(urls))
	defer close(input)

	/* 将要扫描的host放到甬道中 */
	go func(){
		for _, url := range(urls){
			input <- url+ Par.DirInfo.Dir
		}
	}()
	thread := 10
	if len(urls) < Par.Thread {
		thread = len(urls)
	} else {
		thread = Par.Thread
	}
	for i := 0; i< thread; i++{
		go scanDir(input,result)
	}
	public.Out(result,Par.Ofile)
}

func scanDir(input chan string,result chan string){
	for {
		task,ok := <-input
		if !ok{
			return
		}
		//fmt.Println(task)
		if find := strings.Contains(task, "http"); find {
			getDir(task,result)
		} else {
			result<-""
		}
	}
}

func getDir(url string,result chan string){
	/* 跳过https验证 */
	var httpPar = public.NewHttpPar()
	httpPar.Url = url
	httpPar.Timeout = Par.Timeout
	httpPar.Follow = Par.Follow

	header, body, code := public.HttpDoGet2Body(httpPar)
	//fmt.Println(string(body))
	//matched, _ = regexp.MatchString(rule.Rule, Headers)
	if code == 0 {
		result<-""
	} else if code == Par.DirInfo.Code{
		/* 只看code */
		if Par.DirInfo.Header=="" && Par.DirInfo.Body==""{
			result<-url
			/* 同时看code,body,header */
		} else if Par.DirInfo.Header!="" && Par.DirInfo.Body!=""{
			if findbody := strings.Contains(string(body), Par.DirInfo.Body);findbody{
				//if findheader := strings.Contains(fmt.Sprintf("%s",header), Par.DirInfo.Header);findheader{
				if find, _ := regexp.MatchString(Par.DirInfo.Header, fmt.Sprintf("%s",header)); find{
					result<-url
				} else {
					result<-""
				}
			}
		} else {
			/* 判断code和header */
			if Par.DirInfo.Header!=""{
				//if findheader := strings.Contains(fmt.Sprintf("%s",header), Par.DirInfo.Header);findheader{
				if find, _ := regexp.MatchString(Par.DirInfo.Header, fmt.Sprintf("%s",header)); find{
					result<-url
				} else {
					result<-""
				}
				/* 判断code和body */
			} else {
				//if findbody := strings.Contains(string(body), Par.DirInfo.Body);findbody{
				if find, _ := regexp.MatchString(Par.DirInfo.Body, string(body)); find{
					result<-url
				} else {
					result<-""
				}
			}
		}
	} else {
		/* code不等 */
		result<-""
	}
}