# iris-seckill
go商城秒杀系统，使用iris框架开发。

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220708180238.png)

# 技术栈

- Iris
- gorm
- MySQL
- Redis
- RabbitMQ
- go `html/template`包
- 一致性Hash算法

一致性Hash算法：

原理：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220802160050.png)

解决服务器较少情况下的数据倾斜问题：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220802155034.png)



# 路由分析

通过controller如何分析出路由呢，以GET请求，`/product/all`为例：

在main.go的注册路由器中：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220729173824.png)

在`controller/product.go`中：

![](https://img-qingbo.oss-cn-beijing.aliyuncs.com/img/20220729173903.png)

所以这个接口的cURL是：

```
curl --location --request GET 'http://127.0.0.1:8080/product/all'
```













