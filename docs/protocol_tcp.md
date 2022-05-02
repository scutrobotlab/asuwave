# 通信协议

## 示意图

    ```
    +-------------+ Serial +-------------+  TCP   +-------------+
    |     MCU     |------->|   Server    |------->|   Client    |
    |             |<-------|   Backend   |<-------|   Frontend  |
    +-------------+        +-------------+        +-------------+
    ```

1. 如无特殊说明，所有请求和响应参数均为JSON格式。  
2. 请求成功则状态码为2XX。  
3. 请求失败状态码为4XX或5XX，具体错误信息在返回JSON的Error中。  

## 1. 变量

### 1.1 查看订阅变量的值
* 请求地址  

    |  URL  |
    |-------|
    | `/ws` |
* 响应结果  

    |        参数        |     类型     |   说明   |
    |-------------------|--------------|---------|
    | Variables         | array struct | 变量列表 |
    | Variables[].Board | int          | 板子代号 |
    | Variables[].Name  | string       | 变量名   |
    | Variables[].Data  | float        | 变量值   |
    | Variables[].Tick  | int          | 时间戳   |
* 调用示例  

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
