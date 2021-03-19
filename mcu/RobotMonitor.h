/**
 ******************************************************************************
 * @file      RobotMonitor.h
 * @author    M3chD09
 * @brief     Header file of RobotMonitor.c
 * @version   1.1
 * @date      20th Sept 2019
 ******************************************************************************

 ******************************************************************************
 */
#ifndef _ROBOTMONITOR_H_
#define _ROBOTMONITOR_H_

#ifdef __cplusplus
 extern "C" {
#endif

/* Includes ------------------------------------------------------------------*/
#include "string.h"
#include "stm32f4xx_hal.h"
#include "cmsis_os.h"

extern osThreadId_t robotMonitorHandle;

void receiveRobotMonitorRxBuf(void);
void tskRobotMonitor(void *argument);

#ifdef __cplusplus
}
#endif

#endif
