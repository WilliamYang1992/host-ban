# host-ban
一个快速在 hosts 文件添加重定向某些 host 到其它 IP 地址 (如0.0.0.0) 的命令行工具  
  
#### 兼容系统:   
**Windows, MacOS, Linux**  

#### 使用方法:  
**参数**:  
-ip: 转发的 ip 地址, 不指定默认为 `0.0.0.0`  
-host: host 地址, 如 `www.somesite.com`, 不需要带 http(s)://   
-action: 可选值为 `add` 或者 `delete`, 分别代表增加或删除, 不指定默认为 `add`  
  
**增加规则**:  
`host-ban -host www.somesite.com`  或  
`host-ban -ip 0.0.0.0 -host www.somesite.com -action add`
  
**删除规则**:  
`host-ban -host www.somesite.com -action delete`

#### 常用使用场景
1. jexxrains 产品更新与激活快速切换
2. 禁止某些软件请求广告连接