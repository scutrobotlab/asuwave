# 上位机 · 网页版 · 重制版

![Go](https://github.com/scutrobotlab/asuwave/workflows/Go/badge.svg) ![Node.js](https://github.com/scutrobotlab/asuwave/workflows/Node.js/badge.svg) ![Release](https://github.com/scutrobotlab/asuwave/workflows/Release/badge.svg) ![Heroku](https://github.com/scutrobotlab/asuwave/workflows/Heroku/badge.svg)

[Demo](https://asuwave.herokuapp.com/)  

>~~你所用过坠好的上位机~~  
>~~简洁、优雅且好用~~  
>每日一问，今日进度如何

![logo](src/assets/logo.png)

## 佛系·上位机

君问何项最佛系，当属网页上位机。  
先由杨编做设计，后有玮文提建议。  
退队人员再召集，分工明确尽全力。  
年末网管协议拟，初有成效心欢喜。  
世事难料众人离，一昧孤行无所依。  
来年再把项目启，当年锐气远不及。  
半途而废人言弃，奈何图表帧数低。  
次年又把决心立，志在月底创佳绩。  
后端进展颇顺利，前端不见人踪迹。  
五月将至无人理，项目组员心已急。  
两日奋战舞士气，何时完成仍成谜。  
回首往事泪满地，此事羞与后人提。  

## 使用教程

### 单片机端
在 *RobotMonitor.c* 中修改 `RM_UART` 的值为与上位机通讯的串口号，默认为串口1  
在 *stm32f4xx_it.c* 和 *freertos.c* 中添加以下代码  
```C
#include "RobotMonitor.h"
```
在 `USARTx_IRQHandler` 中调用 `receiveRobotMonitorRxBuf()`  
在 *freertos.c* 中的 `USER CODE BEGIN RTOS_THREADS` 与 `USER CODE END RTOS_THREADS` 之间添加以下代码  
```C
const osThreadAttr_t robotMonitor_attributes = {
  .name = "robotMonitor",
  .priority = (osPriority_t) osPriorityNormal,
  .stack_size = 1024
};
robotMonitorHandle = osThreadNew(tskRobotMonitor, NULL, &robotMonitor_attributes);
```

### 上位机网页端
选择并开启串口后，在 *读取变量列表* 中上传elf或者axf文件，在程序成功解析之后就可以添加变量，其中：  

* 观察 只能用于绘制曲线  
* 修改 只能用于调参  

添加完成后，在图表中观察绘制的曲线，在调参面板调参  
