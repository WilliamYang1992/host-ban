# host-ban
一个快速在 hosts 文件添加重定向某些 host 到其它 IP 地址 (如0.0.0.0) 的命令行工具  
  
#### 兼容系统: Windows, Macos, Linux  

#### 使用方法:  
增加规则:  
`host-ban -ip 0.0.0.0 -host www.somesite.com -action add`
  
删除规则:  
`host-ban -host www.somesite.com -action delete`

#### 常用使用场景
1. jexxrains 产品更新与激活快速切换
2. 禁止某些软件请求广告连接