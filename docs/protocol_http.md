# 通信协议

## 前言

    ```
    +-------------+ Serial +-------------+  TCP   +-------------+
    |     MCU     |------->|   Server    |------->|   Client    |
    |             |<-------|   Backend   |<-------|   Frontend  |
    +-------------+        +-------------+        +-------------+
    ```

1. 如无特殊说明，所有请求和响应参数均为JSON格式。  
2. 请求成功则状态码为2XX。  
3. 请求失败状态码为4XX或5XX，具体错误信息在返回JSON的Error中。  

## 1. 串口

### 1.1 获取当前可用的串口列表
* 请求地址  

    |  方法  |    URL    |
    |-------|-----------|
    | `GET` | `/serial` |
* 请求参数  

    无  
* 响应结果  

    |   参数   |     类型     |   说明   |
    |---------|--------------|---------|
    | Serials | array string | 串口列表 |
* 调用示例  

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

### 1.2 获取当前已打开的串口
* 请求地址  

    |  方法  |      URL      |
    |-------|---------------|
    | `GET` | `/serial_cur` |
* 请求参数  

    无  
* 响应结果  

    |  参数   |  类型  |  说明  |
    |--------|--------|--------|
    | Serial | string | 串口名 |
    | Baud   | int    | 波特率 |
    | Status | bool   | 状态  |
* 调用示例  

    请求示例：  
    `GET /serial_cur`  
    响应示例：  
    ```json
    {
        "Serial": "COM3",
        "Baud": 119200,
    }
    ```

### 1.3 打开串口
* 请求地址  

    |  方法  |      URL       |
    |--------|---------------|
    | `POST` | `/serial_cur` |
* 请求参数  

    |  参数   |  类型  |  说明  |
    |--------|--------|-------|
    | Serial | string | 串口名 |
    | Baud   | int    | 波特率 |
* 响应结果  

    无
* 调用示例  

    请求示例：  
    `POST /serial_cur`  
    ```json
    {
        "Serial": "COM3",
        "Baud": 119200
    }
    ```
    响应示例：  
    无

### 1.4 关闭串口
* 请求地址  

    |   方法    |      URL       |
    |----------|----------------|
    | `DELETE` | `/serial_cur`  |
* 请求参数  

    无  
* 响应结果  

    无  
* 调用示例  

    请求示例：  
    `DELETE /serial_cur`  
    响应示例：  
    无  

## 2. 变量

### 2.1 获取支持的变量类型
* 请求地址  

    |  方法  |       URL        |
    |-------|------------------|
    | `GET` | `/variable_type` |
* 请求参数  

    无  
* 响应结果  

    |  参数  |    类型      |    说明     |
    |-------|--------------|------------|
    | Types | array string | 变量类型列表 |
* 调用示例  

    请求示例：  
    `GET /variable_type`  
    响应示例：  
    ```json
    {
        "Types":[
            "double","float","int","int16_t","int32_t","int64_t","int8_t","uint16_t","uint32_t","uint64_t","uint8_t"
        ]
    }
    ```

### 2.2 获取订阅变量
* 请求地址  

    |  方法  |       URL        |
    |-------|------------------|
    | `GET` | `/variable_read` |
* 请求参数  

    无  
* 响应结果  

    |        参数        |     类型     |   说明   |
    |-------------------|--------------|---------|
    | Variables         | array struct | 变量列表 |
    | Variables[].Board | int          | 板子代号 |
    | Variables[].Name  | string       | 变量名   |
    | Variables[].Type  | string       | 变量类型 |
    | Variables[].Addr  | int          | 变量地址 |
    | Variables[].Data  | float        | 变量值   |
* 调用示例  

    请求示例：  
    `GET /variable_read`  
    响应示例：  
    ```json
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
                "Addr":536890180,
                "Data":0
            }
        ]
    }
    ```

### 2.3 添加订阅变量
* 请求地址  

    |  方法  |        URL        |
    |--------|------------------|
    | `POST` | `/variable_read` |
* 请求参数  

    |  参数  | 类型   |   说明   |
    |-------|--------|---------|
    | Board | int    | 板子代号 |
    | Name  | string | 变量名   |
    | Type  | string | 变量类型 |
    | Addr  | int    | 变量地址 |
* 响应结果  

    无  
* 调用示例  

    请求示例：  
    `POST /variable_read`  
    ```json
    {
        "Board":1,
        "Name":"traceme",
        "Type":"float",
        "Addr":536889920
    }
    ```
    响应示例：  
    无  

### 2.4 删除订阅变量
* 请求地址  

    |   方法    |        URL       |
    |----------|------------------|
    | `DELETE` | `/variable_read` |
* 请求参数  

    |  参数  | 类型   |   说明   |
    |-------|--------|---------|
    | Board | int    | 板子代号 |
    | Name  | string | 变量名   |
    | Type  | string | 变量类型 |
    | Addr  | int    | 变量地址 |
* 响应结果  

    无  
* 调用示例  

    请求示例：  
    `DELETE /variable_read`  
    ```json
    {
        "Board":1,
        "Name":"traceme",
        "Type":"float",
        "Addr":536889920
    }
    ```
    响应示例：  
    无  

### 2.5 获取调参变量
* 请求地址  

    |  方法  |       URL        |
    |-------|------------------|
    | `GET` | `/variable_write` |
* 请求参数  

    无  
* 响应结果  

    |        参数        |     类型     |   说明   |
    |-------------------|--------------|---------|
    | Variables         | array struct | 变量列表 |
    | Variables[].Board | int          | 板子代号 |
    | Variables[].Name  | string       | 变量名   |
    | Variables[].Type  | string       | 变量类型 |
    | Variables[].Addr  | int          | 变量地址 |
    | Variables[].Data  | float        | 变量值   |
* 调用示例  

    请求示例：  
    `GET /variable_write`  
    响应示例：  
    ```json
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
                "Addr":536890180,
                "Data":0
            }
        ]
    }
    ```

### 2.6 添加调参变量
* 请求地址  

    |  方法  |        URL        |
    |--------|------------------|
    | `POST` | `/variable_write` |
* 请求参数  

    |  参数  | 类型   |   说明   |
    |-------|--------|---------|
    | Board | int    | 板子代号 |
    | Name  | string | 变量名   |
    | Type  | string | 变量类型 |
    | Addr  | int    | 变量地址 |
* 响应结果  

    无  
* 调用示例  

    请求示例：  
    `POST /variable_write`  
    ```json
    {
        "Board":1,
        "Name":"traceme",
        "Type":"float",
        "Addr":536889920
    }
    ```
    响应示例：  
    无  

### 2.7 修改调参变量的值
* 请求地址  

    |  方法  |        URL       |
    |-------|------------------|
    | `PUT` | `/variable_write` |
* 请求参数  

    |  参数  | 类型   |   说明   |
    |-------|--------|---------|
    | Board | int    | 板子代号 |
    | Name  | string | 变量名   |
    | Type  | string | 变量类型 |
    | Addr  | int    | 变量地址 |
    | Data  | float  | 变量值   |
* 响应结果  

    无  
* 调用示例  

    请求示例：  
    `PUT /variable_write`  
    ```json
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

### 2.8 删除调参变量
* 请求地址  

    |   方法    |        URL       |
    |----------|------------------|
    | `DELETE` | `/variable_write` |
* 请求参数  

    |  参数  | 类型   |   说明   |
    |-------|--------|---------|
    | Board | int    | 板子代号 |
    | Name  | string | 变量名   |
    | Type  | string | 变量类型 |
    | Addr  | int    | 变量地址 |
* 响应结果  

    无  
* 调用示例  

    请求示例：  
    `DELETE /variable_write`  
    ```json
    {
        "Board":1,
        "Name":"traceme",
        "Type":"float",
        "Addr":536889920
    }
    ```
    响应示例：  
    无  

### 2.9 获取工程变量
* 请求地址  

    |  方法  |       URL        |
    |-------|------------------|
    | `GET` | `/variable_proj` |
* 请求参数  

    无  
* 响应结果  

    |        参数        |     类型     |   说明   |
    |-------------------|--------------|---------|
    | Variables         | array struct | 变量列表 |
    | Variables[].Name  | string       | 变量名   |
    | Variables[].Type  | string       | 变量类型 |
    | Variables[].Addr  | int          | 变量地址 |
* 调用示例  

    请求示例：  
    `GET /variable_proj`  
    响应示例：  
    ```json
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
                "Addr":536890180
            }
        ]
    }
    ```

## 3. 工程文件相关

### 3.1 上传工程文件
> 注意：上传工程文件会导致文件监控被清除
* 请求地址  

    |  方法   |       URL        |
    |--------|------------------|
    | `PUT` | `/file/upload`   |
* 请求参数  

    |  参数 | 类型 |     说明      |
    |------|------|--------------|
    | file | 文件 | axf或者elf文件 |
* 响应结果  

    无  
* 调用示例  

    请求示例：  
    `PUT /file/upload`  
    ```
    file=@asuwave.elf
    ```
    响应示例：  
    无  

### 3.2 获取监听的工程文件路径
* 请求地址  

    |  方法  |       URL        |
    |--------|------------------|
    | `GET`  |  `/file/path`   |
* 请求参数  

    无
* 响应结果  

    | 参数 |  类型  |       说明        |
    |------|--------|-------------------|
    | Path | string | axf或者elf文件路径 |
* 调用示例  

    请求示例：  
    `GET /file/path`  
    ```json
    {
        "Path": "C:/user/scutrobotlab/robot.axf"
    }
    ```
    响应示例：  
    无  

### 3.3 设置监听的工程文件路径
* 请求地址  

    |  方法   |       URL        |
    |---------|------------------|
    |  `PUT`  | `/file/path`    |
* 请求参数  

    | 参数  |  类型  |       说明        |
    |------|--------|-------------------|
    | Path | string | axf或者elf文件路径 |
* 响应结果  

    无  
* 调用示例  

    请求示例：  
    `PUT /file/path`  
    ```json
    {
        "Path": "C:/user/scutrobotlab/robot.axf"
    }
    ```
    响应示例：  
    无  

### 3.4 清除监听的工程文件路径
* 请求地址  

    |  方法   |       URL        |
    |--------|------------------|
    | `DELETE` | `/file/path`   |
* 请求参数  

    无
* 响应结果  

    无  
* 调用示例  

    请求示例：  
    无
    响应示例：  
    无  

## 4. 设置

### 4.1 查看设置
* 请求地址  

    |  方法  |    URL     |
    |-------|------------|
    | `GET` | `/option`  |
* 请求参数  

    无
* 响应结果  

    |  参数  |  类型  |   说明    |
    |-------|-------|-----------|
    | Save  |  int  |  保存选项  |
* 调用示例  

    请求示例：  
    `GET /option`  
    响应示例：  
    ```json
    {
        "Save": 7
    }
    ```

### 4.2 修改设置
* 请求地址  

    |  方法  |    URL     |
    |-------|------------|
    | `PUT` | `/option`  |
* 请求参数  

    |  参数  |  类型  |   说明    |
    |-------|-------|-----------|
    | Save  |  int  |  保存选项  |
* 响应结果  

    |  参数  |  类型  |   说明    |
    |-------|-------|-----------|
    | Save  |  int  |  保存选项  |
* 调用示例  

    请求示例：  
    `PUT /option`  
    ```json
    {
        "Save": 6
    }
    ```
    响应示例：  
    ```json
    {
        "Save": 6
    }
    ```
