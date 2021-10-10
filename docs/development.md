# 开发指南

## 电脑端

* node --version >= v14 （建议）  
* go version >= 1.17 （必要）  

### 常用命令
```bash
# 前端
npm ci # 安装依赖
npm run serve # 启动并调试
npm run build # 生产环境构建
# 后端
go test -v ./... # 测试
go build # 编译开发版
./asuwave # 执行
# 构建Release
chmod +x build.sh
./build.sh
```
