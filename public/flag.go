package public

import "flag"

func Flag() {
	var svar string
	Banner()
	// 通用参数
	flag.StringVar(&InputValue.ModePtr,"mode", "all", "运行模式如:webscan、brute、title、fofa、mysql、mssql等")
	flag.StringVar(&InputValue.IpsPtr,"ip","","如:127.0.0.1、127.0.0.1/24、127.0.0.1-255")
	flag.StringVar(&InputValue.IfilePtr,"ifile", "", "输入文件")
	flag.StringVar(&InputValue.OfilePtr,"ofile", "", "输出文件")
	flag.IntVar(&InputValue.ThreadPtr,"thread", 0, "thread")
	flag.IntVar(&InputValue.TimePtr,"time", 0, "timeout")
	flag.BoolVar(&InputValue.HelpPtr,"help", false, "help")
	flag.StringVar(&InputValue.UrlPtr,"url", "", "url")
	flag.BoolVar(&InputValue.NoPingPtr,"np",false,"np，不进行icmp存活检测")

	// 端口扫描参数
	flag.StringVar(&InputValue.PortPtr,"port","","扫描端口如:80,443-445,8000-9000")
	flag.BoolVar(&InputValue.WebPtr,"web",false,"port和fofa host输出格式如:http://127.0.0.1:80")

	// dir扫描参数
	flag.StringVar(&InputValue.DirPtr,"dir", "", "dir如:/login.jsp")
	flag.IntVar(&InputValue.CodePtr,"code", 200, "dir返回code如:200、302")
	flag.StringVar(&InputValue.HeaderPtr,"header", "", "dir返回header如:rememberMe")
	flag.StringVar(&InputValue.BodyPtr,"body", "", "dir返回body如:root:x:0:0")

	// fofa
	flag.StringVar(&InputValue.KeyPtr,"key", "", "fofa查询语句如:domain='fofa.so'")
	flag.IntVar(&InputValue.NumPtr,"num", 0, "fofa请求数量如:100、10000")
	flag.StringVar(&InputValue.FieldsPtr,"fields", "", "fofa返回类型如:ip,port")

	// brute
	flag.StringVar(&InputValue.UserFilePtr,"ufile", "", "用户字典")
	flag.StringVar(&InputValue.PassFilePtr,"pfile", "", "密码字典")
	flag.StringVar(&InputValue.UserPtr,"user", "", "用户")
	flag.StringVar(&InputValue.PassPtr,"pass", "", "密码")

	flag.StringVar(&svar, "svar", "bar", "a string var") // 对变量取址
	flag.Parse()
}