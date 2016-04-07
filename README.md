### gorest说明

## 设计功能

> * 使用Golang把PostgreSQL的CRUD接口封装为RESTful服务
> * gorest.conf中加载配置（可以同步修改？）
> * 在配置文件中配置数据库服务、跨域调用以及其他（支持多数据库？）
> * 使用json格式输入输出（也可支持多种方式？）
> * 可以做到可配置路由？

## json格式

``` javascript
[
	{ id: "id1", columnName1: "value1", columnName2: "value2" },
]
```

参考：https://github.com/emicklei/mora (完全没参考)

## 完成度

> * 完成了postgresql和mysql的接入
> * 完成了GET/POST/PUT/DELETE四个动词对单个表的操作
> * 未完成复杂情况下的测试
> * 未完成status code的设置
> * 未完成关联表的查询
> * 未完成名词命名规则的转换
> * 未完成返回资源的url# gorest
