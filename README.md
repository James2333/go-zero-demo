# 场景
程序员小明需要借阅一本《西游记》，在没有线上图书管理系统的时候，他每天都要去图书馆前台咨询图书馆管理员，
* 小明：你好，请问今天《西游记》的图书还有吗？
* 管理员：没有了，明天再来看看吧。

过了一天，小明又来到图书馆，问：
* 小明：你好，请问今天《西游记》的图书还有吗？
* 管理员：没有了，你过两天再来看看吧。

就这样经过多次反复，小明也是徒劳无功，浪费大量时间在来回的路上，于是终于忍受不了落后的图书管理系统，
他决定自己亲手做一个图书查阅系统。

## 预期实现目标
* 用户登录
  依靠现有学生系统数据进行登录
* 图书检索
  根据图书关键字搜索图书，查询图书借阅情况，归还时间等。

## 系统分析
### 服务拆分
* user
    * api 提供用户登录协议
    * rpc 供search服务访问用户数据
* search
    * api 提供图书查询协议

> [!TIP]
> 这个微小的图书借阅查询系统虽然小，从实际来讲不太符合业务场景，但是仅上面两个功能，已经满足我们对go-zeroapi/rpc的场景演示了，
> 后续为了满足更丰富的go-zero功能演示，会在文档中进行业务插入即相关功能描述。这里仅用一个场景进行引入。


# 部署文档
### model文件生成
不得不说这里真的帮我们开发省去很多的麻烦，只需要sql文件就能生成相应的crud代码
有需要的话可以直接copy二次开发,这里进到/user/model文件下
```
goctl model mysql ddl -src user.sql -dir . -c
```

### 配置文件修改
记得把三个服务etc/*.yaml文件里的redis,mysql,etcd改成自己的服务器地址
并启动这些服务
```
//启动service服务
cd /service/search/api 
go run search.go -f etc/search-api.yaml
//启动user服务
cd /user/api
go run user.go -f etc/user-api.yaml
//启动user的rpc服务
cd /user/rpc
go run user.go  -f etc/user.yaml
```

之后就可以用curl命令测试了