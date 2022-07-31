更新
1、 dir扫描支持正则
3、 修复finger一些bug
4、 修复-ifile一些bug

删除
- 备份扫描
- 指纹库去掉防火墙

日志打印：
wsarecv: An existing connection was forcibly closed by the remote host.
解决方案：
github.com/go-sql-driver/mysql/packets.go 中注释 errLog.Print(err) 

日志打印：
Unsolicited response received on idle HTTP channel starting with
解决方案：
net/http/transport.go 中注释 log.Printf("Unsolicited response received on idle HTTP channel starting with %q; err=%v", buf, peekErr)

