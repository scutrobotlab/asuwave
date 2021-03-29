# 目录结构

```
├── .github - github actions相关
│   └── workflows
│       ├── go.yml
│       ├── heroku.yml
│       ├── node.yml
│       └── release.yml
├── datautil - 后端·协议处理
│   ├── const.go
│   ├── datautil.go
│   └── datautil_test.go
├── fromelf - 后端·elf/axf文件处理
│   └── fromelf.go
├── mcu  - 单片机端
│   ├── README.md
│   ├── RobotMonitor.c
│   └── RobotMonitor.h
├── option -  后端·用户配置控制
│   ├── jsonfile.go
│   ├── option.go
│   └── option_test.go
├── public - 前端·假入口
│   ├── favicon.ico
│   └── index.html
├── serial - 后端·串口控制
│   ├── serial.go
│   └── testport.go
├── server - 后端·网页接口
│   ├── ctrler_test.go
│   ├── optionctrl.go
│   ├── optionctrl_test.go
│   ├── serialctrl.go
│   ├── serialctrl_test.go
│   ├── server.go
│   ├── variablectrl.go
│   ├── variablectrl_test.go
│   └── websocket.go
├── src - 前端·真入口
│   ├── api
│   │   ├── internal.js
│   │   ├── option.js
│   │   ├── serial.js
│   │   └── variable.js
│   ├── assets
│   │   └── logo.png
│   ├── components
│   │   ├── AboutDialog.vue
│   │   ├── ChartCard.vue
│   │   ├── DrawerList.vue
│   │   ├── ErrorAlert.vue
│   │   ├── PanelCard.vue
│   │   ├── SaveConfig.vue
│   │   ├── SerialPort.vue
│   │   ├── VariableAllDialog.vue
│   │   ├── VariableList.vue
│   │   └── VariableNewDialog.vue
│   ├── mixins
│   │   └── errorMixin.js
│   ├── plugins
│   │   └── vuetify.js
│   ├── store
│   │   └── index.js
│   ├── App.vue
│   └── main.js
├── variable - 后端·变量
│   ├── typeconvert.go
│   ├── typeconvert_test.go
│   └── variable.go
├── .editorconfig
├── .gitignore
├── .prettierrc - 自定义前端文件格式化规则
├── babel.config.js
├── build.sh - 用于构建Release的脚本
├── Dockerfile - 用于部署演示网页到Heroku的docker
├── go.mod
├── go.sum
├── LICENSE
├── main.go - 后端入口
├── package-lock.json
├── package.json
├── README.md
└── vue.config.js
```
