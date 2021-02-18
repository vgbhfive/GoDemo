## shortLink
使用 Go 语言实现的短网址服务

### 接口
 - 创建短链接口： 
```
POST http://127.0.0.1:9000/api/shorten
{
   	"url": "http://www.hiningmeng.cn",
   	"expiration_in_minutes": 30
}
```

 - 短链详细信息
```
GET http://127.0.0.1:9000/api/info?shortlink=1
```

 - 短链跳转真实长链地址
```
GET http://127.0.0.1:9000/1
```