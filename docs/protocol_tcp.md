# 通信协议

## 示意图

```
+-------------+ Serial +-------------+  TCP   +-------------+
|     MCU     |------->|   Server    |------->|   Client    |
|             |<-------|   Backend   |<-------|   Frontend  |
+-------------+        +-------------+        +-------------+
```

## TCP

### HTTP接口

* 如无特殊说明，所有请求和响应参数均为JSON格式。  
* 请求成功则状态码为2XX。  
* 请求失败状态码为4XX或5XX，具体错误信息在返回JSON的Error中。  

#### 串口

##### 获取当前可用的串口列表
###### 请求地址
|  方法  |    URL    |
|-------|-----------|
| `GET` | `/serial` |
###### 请求参数
无  
###### 响应结果
|   参数   |     类型     |   说明   |
|---------|--------------|---------|
| Serials | array string | 串口列表 |
###### 调用示例
请求示例：  
`GET /serial`  
响应示例：  
```
{
    "Serials":[
        "COM3",
        "COM4"
    ]
}
```

##### 获取当前已打开的串口
###### 请求地址
|  方法  |      URL      |
|-------|---------------|
| `GET` | `/serial_cur` |
###### 请求参数
无  
###### 响应结果
|  参数   |  类型  |  说明  |
|--------|--------|-------|
| Serial | string | 串口名 |
###### 调用示例
请求示例：  
`GET /serial_cur`  
响应示例：  
```
{
    "Serial": "COM3"
}
```

##### 打开串口
###### 请求地址
|  方法  |      URL       |
|--------|---------------|
| `POST` | `/serial_cur` |
###### 请求参数
|  参数   |  类型  |  说明  |
|--------|--------|-------|
| Serial | string | 串口名 |
###### 响应结果
|  参数   |  类型  |  说明  |
|--------|--------|-------|
| Serial | string | 串口名 |
###### 调用示例
请求示例：  
`POST /serial_cur`  
```
{
    "Serial": "COM3"
}
```
响应示例：  
```
{
    "Serial": "COM3"
}
```

##### 关闭串口
###### 请求地址
|   方法    |      URL       |
|----------|----------------|
| `DELETE` | `/serial_cur`  |
###### 请求参数
无  
###### 响应结果
无  
###### 调用示例
请求示例：  
`DELETE /serial_cur`  
响应示例：  
无  

#### 变量

##### 获取支持的变量类型
###### 请求地址
|  方法  |       URL        |
|-------|------------------|
| `GET` | `/variable_type` |
###### 请求参数
无  
###### 响应结果
|  参数  |    类型      |    说明     |
|-------|--------------|------------|
| Types | array string | 变量类型列表 |
###### 调用示例
请求示例：  
`GET /variable_type`  
响应示例：  
```
{
    "Types":[
        "double","float","int","int16_t","int32_t","int64_t","int8_t","uint16_t","uint32_t","uint64_t","uint8_t"
    ]
}
```

##### 获取订阅变量
###### 请求地址
|  方法  |       URL        |
|-------|------------------|
| `GET` | `/variable_read` |
###### 请求参数
无  
###### 响应结果
|        参数        |     类型     |   说明   |
|-------------------|--------------|---------|
| Variables         | array struct | 变量列表 |
| Variables[].Board | int          | 板子代号 |
| Variables[].Name  | string       | 变量名   |
| Variables[].Type  | string       | 变量类型 |
| Variables[].Addr  | int          | 变量地址 |
| Variables[].Data  | float        | 变量值   |
###### 调用示例
请求示例：  
`GET /variable_read`  
响应示例：  
```
{
    "Variables":[
        {
            "Board":1,
            "Name":"traceme",
            "Type":"float",
            "Addr":536889920,
            "Data":0
        },
        {
            "Board":1,
            "Name":"count",
            "Type":"int",
            "Addr":‭536890180‬,
            "Data":0
        }
    ]
}
```

##### 添加订阅变量
###### 请求地址
|  方法  |        URL        |
|--------|------------------|
| `POST` | `/variable_read` |
###### 请求参数
|  参数  | 类型   |   说明   |
|-------|--------|---------|
| Board | int    | 板子代号 |
| Name  | string | 变量名   |
| Type  | string | 变量类型 |
| Addr  | int    | 变量地址 |
###### 响应结果
无  
###### 调用示例
请求示例：  
`POST /variable_read`  
```
{
    "Board":1,
    "Name":"traceme",
    "Type":"float",
    "Addr":536889920
}
```
响应示例：  
无  

##### 删除订阅变量
###### 请求地址
|   方法    |        URL       |
|----------|------------------|
| `DELETE` | `/variable_read` |
###### 请求参数
|  参数  | 类型   |   说明   |
|-------|--------|---------|
| Board | int    | 板子代号 |
| Name  | string | 变量名   |
| Type  | string | 变量类型 |
| Addr  | int    | 变量地址 |
###### 响应结果
无  
###### 调用示例
请求示例：  
`DELETE /variable_read`  
```
{
    "Board":1,
    "Name":"traceme",
    "Type":"float",
    "Addr":536889920
}
```
响应示例：  
无  

##### 获取调参变量
###### 请求地址
|  方法  |       URL        |
|-------|------------------|
| `GET` | `/variable_modi` |
###### 请求参数
无  
###### 响应结果
|        参数        |     类型     |   说明   |
|-------------------|--------------|---------|
| Variables         | array struct | 变量列表 |
| Variables[].Board | int          | 板子代号 |
| Variables[].Name  | string       | 变量名   |
| Variables[].Type  | string       | 变量类型 |
| Variables[].Addr  | int          | 变量地址 |
| Variables[].Data  | float        | 变量值   |
###### 调用示例
请求示例：  
`GET /variable_modi`  
响应示例：  
```
{
    "Variables":[
        {
            "Board":1,
            "Name":"traceme",
            "Type":"float",
            "Addr":536889920,
            "Data":0
        },
        {
            "Board":1,
            "Name":"count",
            "Type":"int",
            "Addr":‭536890180‬,
            "Data":0
        }
    ]
}
```

##### 添加调参变量
###### 请求地址
|  方法  |        URL        |
|--------|------------------|
| `POST` | `/variable_modi` |
###### 请求参数
|  参数  | 类型   |   说明   |
|-------|--------|---------|
| Board | int    | 板子代号 |
| Name  | string | 变量名   |
| Type  | string | 变量类型 |
| Addr  | int    | 变量地址 |
###### 响应结果
无  
###### 调用示例
请求示例：  
`POST /variable_modi`  
```
{
    "Board":1,
    "Name":"traceme",
    "Type":"float",
    "Addr":536889920
}
```
响应示例：  
无  

##### 修改调参变量的值
###### 请求地址
|  方法  |        URL       |
|-------|------------------|
| `PUT` | `/variable_modi` |
###### 请求参数
|  参数  | 类型   |   说明   |
|-------|--------|---------|
| Board | int    | 板子代号 |
| Name  | string | 变量名   |
| Type  | string | 变量类型 |
| Addr  | int    | 变量地址 |
| Data  | float  | 变量值   |
###### 响应结果
无  
###### 调用示例
请求示例：  
`PUT /variable_modi`  
```
{
    "Board":1,
    "Name":"traceme",
    "Type":"float",
    "Addr":536889920,
    "Data":1.5
}
```
响应示例：  
无  

##### 删除调参变量
###### 请求地址
|   方法    |        URL       |
|----------|------------------|
| `DELETE` | `/variable_modi` |
###### 请求参数
|  参数  | 类型   |   说明   |
|-------|--------|---------|
| Board | int    | 板子代号 |
| Name  | string | 变量名   |
| Type  | string | 变量类型 |
| Addr  | int    | 变量地址 |
###### 响应结果
无  
###### 调用示例
请求示例：  
`DELETE /variable_modi`  
```
{
    "Board":1,
    "Name":"traceme",
    "Type":"float",
    "Addr":536889920
}
```
响应示例：  
无  

##### 获取工程变量
###### 请求地址
|  方法  |       URL        |
|-------|------------------|
| `GET` | `/variable_proj` |
###### 请求参数
无  
###### 响应结果
|        参数        |     类型     |   说明   |
|-------------------|--------------|---------|
| Variables         | array struct | 变量列表 |
| Variables[].Name  | string       | 变量名   |
| Variables[].Type  | string       | 变量类型 |
| Variables[].Addr  | int          | 变量地址 |
###### 调用示例
请求示例：  
`GET /variable_proj`  
响应示例：  
```
{
    "Variables":[
        {
            "Name":"traceme",
            "Type":"float",
            "Addr":536889920
        },
        {
            "Name":"count",
            "Type":"int",
            "Addr":‭536890180‬
        }
    ]
}
```

##### 上传工程变量
###### 请求地址
|  方法   |       URL        |
|--------|------------------|
| `POST` | `/variable_proj` |
###### 请求参数
|  参数 | 类型 |     说明      |
|------|------|--------------|
| file | 文件 | axf或者elf文件 |
###### 响应结果
无  
###### 调用示例
请求示例：  
`POST /variable_proj`  
```
file=@file.axf
```
响应示例：  
无  

#### 设置

##### 查看设置
###### 请求地址
|  方法  |    URL     |
|-------|------------|
| `GET` | `/option`  |
###### 请求参数
无
###### 响应结果
|  参数  |  类型  |   说明    |
|-------|-------|-----------|
| Save  |  int  |  保存选项  |
###### 调用示例
请求示例：  
`GET /option`  
响应示例：  
```
{
    "Save": 7
}
```

##### 修改设置
###### 请求地址
|  方法  |    URL     |
|-------|------------|
| `PUT` | `/option`  |
###### 请求参数
|  参数  |  类型  |   说明    |
|-------|-------|-----------|
| Save  |  int  |  保存选项  |
###### 响应结果
|  参数  |  类型  |   说明    |
|-------|-------|-----------|
| Save  |  int  |  保存选项  |
###### 调用示例
请求示例：  
`PUT /option`  
```
{
    "Save": 6
}
```
响应示例：  
```
{
    "Save": 6
}
```

### websockect接口

#### 变量

##### 查看订阅变量的值
###### 请求地址
|  URL  |
|-------|
| `/ws` |
###### 响应结果
|        参数        |     类型     |   说明   |
|-------------------|--------------|---------|
| Variables         | array struct | 变量列表 |
| Variables[].Board | int          | 板子代号 |
| Variables[].Name  | string       | 变量名   |
| Variables[].Data  | float        | 变量值   |
| Variables[].Tick  | int          | 时间戳   |
响应示例：  
```
{
    "Variables":[
        {
            "Board":1,
            "Name":"traceme",
            "Data":2.5,
            "Tick":10
        },
        {
            "Board":1,
            "Name":"count",
            "Data":1,
            "Tick":10
        }
    ]
}
```
