[TOC]

##### 简要描述

- SQL查询接口

##### 请求URL

- ` http://ENDPOINT:PORT/sql/query `

##### 请求方式

- POST

##### 参数

| 参数名        | 必填  | 类型     | 说明         |
|:-----------|:----|:-------|------------|
| project    | 是   | int    | 项目id       |
| iid        | 是   | string | 数据源id      |
| sql        | 是   | string | sql        |
| owner      | 是   | string | Database名称 |
| schema     | 否   | string | Schema名称   |
| page_index | 是   | int    | 页码(从0开始)   |
| page_size  | 是   | int    | 每页数量       |
| need_total | 否   | bool   | 是否返回总数     |

##### 返回示例

``` 
{
   "result":1,
   "data":{
      "data":[
         {
            "id":1,
            "sample":"abc"
         }
      ],
      "total":1
   },
   "msg":"ok"
}
```

##### 返回字段说明

| 字段名   | 类型     | 说明   |
|:------|:-------|------|
| data  | object | 查询结果 |
| total | int    | 总数   |

##### 备注
