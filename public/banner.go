package public

import (
	"fmt"
)

func Banner() {
	var goon = ""+
		"---------------------------------------------------------------------------------------------\n" +
		"				  ____  ____  ____  ____ \n"+
		"				 / __ `/ __ \\/ __ \\/ __ \\\n"+
		"				/ /_/ / /_/ / /_/ / / / /\n"+
		"				\\__, /\\____/\\____/_/ /_/\n"+
		"				/___/\n\n"+
		`
					goon v3
					by:i11us0ry
---------------------------------------------------------------------------------------------                                                    
	`
	fmt.Println(goon)
}



func OutMode(){
	var outMode = `
可选mode如下:

all:		默认选项,包含ip-port(web)-title-finger-ftp-ms17010-mssql-mysql-postgres-redis-ssh-smb-rdp-telnet-netbios
webscan:	包含ip-port(web)-title-finger
brute:		包含ip-ftp-ms17010-mssql-mysql-postgres-redis-ssh-smb-rdp-telnet
ip:		ip探活,执行-np可绕过探活
port:		端口扫描,执行-web直接探测http/https
fofa:		fofa资产获取,执行-web输出host时添加http(fields为多个时host放在最后一位)
title:		title
finger:		web指纹
dir:		单路径扫描
back:		压缩文件扫描
tomcat:		tomcat爆破，目标为url，如http://127.0.0.1:8080或http://127.0.0.1:8080/manager/html
ftp:		ftp爆破,其他ms17010,mssql,mysql,postgres,redis,ssh,smb,rdp,telnet同理
netbios:	netbios探测

详细参考：-h

`
	fmt.Println(outMode)
}