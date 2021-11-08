# 2021-06-28 goon_beta_v1
> 1、完成portscan、titlescan、dirscan、fofascan、pluginscan、autoscan功能

# 2021-06-30 goon_beta_v1.1
> 1、删除fofa自动获资产功能  
> 2、可在配置文件中自定义扫描超时时间  
> 3、可选择对fofa返回的host自动添加"http://"  
> 4、fofa资产获取结果自动去重  

# 2021-06-30 goon_beta_v1.2

> 1、 优化 titlescan，如下：
>> 1.1、 多种编码问题  
>> 1.2、 去除 title 前后空格、换行符  
>> 1.3、 首页 js 设置 title  
>> 1.4、 首页各类跳转  
>> 1.5、 纠正正则语法

> 2、 解决了一些小 bug

# 2021-07-06 goon_beta_v1.3
> 1、 优化 titlescan  
> 2、 解决了一些小 bug

# 2021-07-22 goon_beta_v2.0
> 1、 新增主机存活探测(ipscan)  
> 2、 新增备份文件扫描(backscan)  
> 3、 新增爆破模块(brute),包括ssh、redis、mysql、mssql、ftp及postgres爆破  
> 4、 修改了pluginscan的匹配机制，新增了一批内网渗透常见指纹  

# 2021-07-25 goon_beta_v2.1
> 1、 解决titlescan的小bug  
> 2、 新增waf指纹  

# 2021-08-13 goon_beta_v2.2
> 1、修复了众多bug
> 2、重构了框架

# 2021-11-08 goon_beta_v2.3
> 1. 修改port扫描中的web扫描方式，现在如果http存活则放弃探测https是否存活，且可根据配置文件中的状态码返回web资产，如返回常见的200、302、401。
> 2. 爆破模块默认端口可以配置多个，如ssh爆破的默认端口，port: [22, 2222, 22222]
> 3. ftp爆破返回第一个目录名称，减少验证时间
> 4. 爆破模块的待爆破资产同时支持127.0.0.1:22和127.0.0.1两种形式，如果待爆破资产没有指定端口，则会自动添加配置文件中的默认端口进行指纹识别
> 5. 新增awvs联动，搭建好awvs后在配置文件中设置url、key即可
