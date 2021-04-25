# Go作业（Kratos Project Template）

## 第二期作业
```
使用gorm数据库连接
1.在internal/data/data.go中修改数据库连接
2.在api/helloworld/v1/greeter_http.pb.go中增加/get/{name} GET请求
3.在网页上使用get/{name}请求时，检查{name}，如果有数据返回content值，如果没有查询不到name，在业务接口中返回gorm的ErrRecordNotFound，并处理这个指定错误，并返回用户未找到信息给调用端

这样处理的几点思考：
1.不同的API，在通过数据库查询接口获取数据时，有可能只需要获取一个字段。如果不抛异常就需要做数据的封装，判断是不是因为查询不到还是数据库保存的就是空值。
2.对API方便处理，如果没有数据，直接去数据库查询，查询到保存到cache。如果查询不到可以缓存一个nil，可以通过设置一个简单的短期过期策略，让API可以再次获取数据库。
```
