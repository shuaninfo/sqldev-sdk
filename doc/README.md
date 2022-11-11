接口调用说明
======

本文介绍了SQLDEV接口的调用方式及接口文档的相关格式说明。

请求协议
----

HTTP/HTTPS

请求方法
----
GET/POST 参数分别位于query和formdata中

返回示例
----

返回结果为application/json格式

普通返回结构如下：

```json
{
  "result": 1,
  "data": {
    "id": 1,
    "created_at": "2022-11-11T15:12:28+08:00",
    "updated_at": "2022-11-11T15:12:28+08:00",
    "name": "Antares"
  },
  "msg": "ok"
}
```

| 字段名    | 类型     | 说明             |
|:-------|:-------|:---------------|
| result | int    | 1表示成功 其他为错误代码  |
| data   | object | 返回的数据          |
| msg    | string | ok表示成功，其他为错误原因 |

列表返回结构如下

```json
{
  "result": 1,
  "data": {
    "count": 12,
    "page_no": 0,
    "page_size": 100,
    "list": [
      {
        "id": 1,
        "created_at": "2022-11-11T15:12:28+08:00",
        "updated_at": "2022-11-11T15:12:28+08:00",
        "name": "Telesto"
      },
      {
        "id": 2,
        "created_at": "2022-11-12T15:12:28+08:00",
        "updated_at": "2022-11-12T15:12:28+08:00",
        "name": "Clovis"
      }
    ]
  },
  "msg": "ok"
}
```

| 字段名       | 类型    | 说明          |
|:----------|:------|:------------|
| page_no   | int   | 页码，从0开始     |
| page_size | int   | 每页数量        |
| count     | int   | 总数          |
| list      | array | 数据          |

密钥获取
------

如下图所示，在调用SQLDEV接口前，您需要完成以下准备工作：

1、拿到可登录SQLDEV的用户名和密码， 并登陆SQLDEV

2、获取应用标识（app_key）和应用密钥（app_secret）。app_key相当于是身份凭证，调用接口时，在参数里添加app_key以及签名（sign）来鉴权调用者身份。

![应用标识和应用密钥的获取方式](https://www.showdoc.com.cn/server/api/attachment/visitFile?sign=09230fff27f17fa877db6a9198add027 "[Snipaste_2022-11-10_14-38-26.png")

3、签名的说明
所有接口请求时都需要传递app_key和current_time和sign作为参数，其中current_time为毫秒级的unix时间戳。需要将包含app_key和current_time在内的所有参数进行签名运算，将运算结果作为参数sign的值。具体的签名算法可以在SDK内找到，也可以直接使用SDK进行接口调用。


文档格式约定
------

每个接口文档会按照如下格式提供信息：

1. 简要描述
2. 请求URL
3. 请求方式
4. 参数
5. 返回示例
6. 备注

注意事项
------

- 当使用POST方法调用接口时，需设置Content-Type为application/x-www-form-urlencoded


- 所有接口在调用时都会返回result、msg和data，其中result=1时说明该接口调用成功，其他情况则说明该接口业务调用异常，开发者可根据result和msg排查问题。


- 请不要仅根据msg判断调用是否成功。当请求返回结果中返回了result且不为1时可判断为请求失败。msg是对result的说明，供开发者参考排查问题。
