# Robot Monitor Handler
## The best upper monitor you have ever used.
>simple, elegant and useful

### Feature
* No need to add variables in the program, just include it and call functions.
* View and modify variables.

### Usage
Change the value of RM_UART in file *RobotMonitor.c* to the UART used to communicate with NUC. The default value is huart1.  
Add the following code in file *stm32f4xx_it.c* and *freertos.c*.  
```C
#include "RobotMonitor.h"
```
Call `receiveRobotMonitorRxBuf()` in function `USARTx_IRQHandler`.  
Add the following code between "USER CODE BEGIN RTOS_THREADS" and "USER CODE END RTOS_THREADS" in file *freertos.c*.
```C
const osThreadAttr_t robotMonitor_attributes = {
  .name = "robotMonitor",
  .priority = (osPriority_t) osPriorityNormal,
  .stack_size = 1024
};
robotMonitorHandle = osThreadNew(tskRobotMonitor, NULL, &robotMonitor_attributes);
```
