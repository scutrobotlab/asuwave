/**
 ******************************************************************************
 * @file      RobotMonitor.c
 * @author    M3chD09
 * @brief     Robot monitor handler, communicating with NUC
 * @version   1.1
 * @date      20th Sept 2019
 ******************************************************************************

 ==============================================================================
 ##### How to use this tool #####
 ==============================================================================
 [..]
 Change the value of RM_UART in file "RobotMonitor.c" to the UART used to communicate with NUC. The default value is huart1.
 [..]
 Add the following code in file *stm32f4xx_it.c* and *freertos.c*.
 ***CODE BEGIN***
 #include "RobotMonitor.h"
 ***CODE END***
 [..]
 Call "receiveRobotMonitorRxBuf()" in function "USARTx_IRQHandler".
 [..]
 Add the following code between "USER CODE BEGIN RTOS_THREADS" and "USER CODE END RTOS_THREADS" in "freertos.c".
 ***CODE BEGIN***
 const osThreadAttr_t robotMonitor_attributes = {
 .name = "robotMonitor",
 .priority = (osPriority_t) osPriorityNormal,
 .stack_size = 1024
 };
 robotMonitorHandle = osThreadNew(tskRobotMonitor, NULL, &robotMonitor_attributes);
 ***CODE END***


 ******************************************************************************
 */
#include "RobotMonitor.h"

/* UART define */
/* Change huart1 to the UART you use to communicate with the NUC */
extern UART_HandleTypeDef huart1;
#define RM_UART huart1

/**
 * @brief  RM_board_define board ID definition
 */
enum RM_BOARD
{
  RM_BOARD_1 = 0x01,
  RM_BOARD_2,
  RM_BOARD_3
};

/**
 * @brief  action information definition
 */
enum RM_ACT
{
  RM_ACT_SUBSCRIBE = 0x01,
  RM_ACT_SUBSCRIBERETURN,
  RM_ACT_UNSUBSCRIBE,
  RM_ACT_UNSUBSCRIBERETURN,
  RM_ACT_READ,
  RM_ACT_READRETURN,
  RM_ACT_WRITE,
  RM_ACT_WRITERETURN
};

/**
 * @brief  error information definition
 */
enum RM_ERROR
{
  RM_ERROR_NOSUCHADDRREG = 0xf9,
  RM_ERROR_FULLADDR,
  RM_ERROR_NOSUCHDATANUM,
  RM_ERROR_NOSUCHADDR,
  RM_ERROR_NOSUCHACT,
  RM_ERROR_NOSUCHBOARD
};

/** 
 * @brief robot monitor data to receive structure definition
 */
typedef struct
{
  uint8_t board :8; /*!< The ID of the board to be operated.
   This parameter can be any value of @ref RM_BOARD */

  uint8_t act :8; /*!< The action to be performed or the error information to be sent.
   This parameter can be any value of @ref RM_ACT or @ref RM_ERROR */

  uint8_t dataNum :8; /*!< Size amount of data to be operated.
   This parameter must be a number between Min_Data = 0 and Max_Data = 8. */

  uint32_t addr :32; /*!< The address of MCU to be operated.
   This parameter must be a number between Min_Data = 0x20000000 and Max_Data = 0x80000000. */

  uint64_t dataBuf :64; /*!< The data read or to be written.
   This parameter must be a variable of type uint64_t */

  uint8_t carriageReturn :8; /*!< The carriage return.
   This parameter must be '\n' */

}__attribute__((packed)) robotMonitorRxData_t;

/**
 * @brief robot monitor data to transmit structure definition
 */
typedef struct
{
  uint8_t board :8; /*!< The ID of the board to be operated.
   This parameter can be any value of @ref RM_BOARD */

  uint8_t act :8; /*!< The action to be performed or the error information to be sent.
   This parameter can be any value of @ref RM_ACT or @ref RM_ERROR */

  uint8_t dataNum :8; /*!< Size amount of data to be operated.
   This parameter must be a number between Min_Data = 0 and Max_Data = 8. */

  uint32_t addr :32; /*!< The address of MCU to be operated.
   This parameter must be a number between Min_Data = 0x20000000 and Max_Data = 0x80000000. */

  uint64_t dataBuf :64; /*!< The data read or to be written.
   This parameter must be a variable of type uint64_t */

  uint32_t tick :32; /*!< The tick count.
   This parameter must be a variable of type uint32_t */

  uint8_t carriageReturn :8; /*!< The carriage return.
   This parameter must be '\n' */

}__attribute__((packed)) robotMonitorTxData_t;

/** 
 * @brief robot monitor data to receive union definition
 */
typedef union
{
  robotMonitorRxData_t robotMonitorRxData;
  uint8_t robotMonitorRxBuf[16];
} robotMonitorRxUnion_t;
robotMonitorRxUnion_t robotMonitorRxUnion;

/** 
 * @brief robot monitor data to transmit structure definition
 */
typedef union
{
  robotMonitorTxData_t robotMonitorTxData;
  uint8_t robotMonitorTxBuf[20];
} robotMonitorTxUnion_t;
robotMonitorTxUnion_t robotMonitorTxUnion;

/* definition of thread */
osThreadId_t robotMonitorHandle;
/* the flag of the message received form NUC */
uint8_t isGetRobotMonitor;
/* definition of the list of address to read */
#define MAX_ADDR_NUM 5
typedef struct
{
  uint8_t dataNum;
  uint32_t addr;
} listAddr_t;
listAddr_t listAddr[MAX_ADDR_NUM];
/* definition of the data pack and length to send */
static uint8_t txDataBuf[MAX_ADDR_NUM * 20] = {0};
static uint16_t txDataLen = 0;

/**
 * @brief  Receives robot monitor data in idle interrupt mode.
 * @param  None
 * @retval None
 */
void receiveRobotMonitorRxBuf(void)
{
  uint32_t temp = temp;

  /* Check if it is an idle interrupt */
  if ((__HAL_UART_GET_FLAG(&RM_UART,UART_FLAG_IDLE) != RESET))
  {
    isGetRobotMonitor = 1;
    __HAL_UART_CLEAR_IDLEFLAG(&RM_UART);
    HAL_UART_DMAStop(&RM_UART);
    temp = RM_UART.hdmarx->Instance->NDTR;

    /* Receive robot monitor data */
    HAL_UART_Receive_DMA(&RM_UART, robotMonitorRxUnion.robotMonitorRxBuf, sizeof(robotMonitorRxUnion.robotMonitorRxBuf));
  }
}

/**
 * @brief  Regist the address.
 * @param  None
 * @retval the index of the listAddr to regist.
 */
int8_t addrRegister(void)
{
  for (int i = 0; i < MAX_ADDR_NUM; i++)
  {
    /* Find the index of listAddr that is null */
    if (listAddr[i].dataNum == 0 || listAddr[i].addr == robotMonitorRxUnion.robotMonitorRxData.addr)
    {
      listAddr[i].dataNum = robotMonitorRxUnion.robotMonitorRxData.dataNum;
      listAddr[i].addr = robotMonitorRxUnion.robotMonitorRxData.addr;
      return i;
    }
  }
  return -1;
}

/**
 * @brief  Unregist the address.
 * @param  None
 * @retval the index of the listAddr to unregist.
 */
int8_t addrUnregister(void)
{
  for (int i = 0; i < MAX_ADDR_NUM; i++)
  {
    /* Find the index of listAddr */
    if (listAddr[i].addr == robotMonitorRxUnion.robotMonitorRxData.addr)
    {
      listAddr[i].dataNum = 0;
      listAddr[i].addr = 0;
      return i;
    }
  }
  return -1;
}

/**
 * @brief  Send an error message to the NUC.
 * @param  err: the error information.
 * @retval None
 */
void returnError(uint8_t err)
{
  /* Clear the data buffer to be sent */
  memset((uint8_t *) &robotMonitorTxUnion.robotMonitorTxBuf, 0, sizeof(robotMonitorTxUnion.robotMonitorTxBuf));

  /* Prepare the data to send */
  robotMonitorTxUnion.robotMonitorTxData.board = robotMonitorRxUnion.robotMonitorRxData.board;
  robotMonitorTxUnion.robotMonitorTxData.act = err;
  robotMonitorTxUnion.robotMonitorTxData.addr = robotMonitorRxUnion.robotMonitorRxData.addr;
  robotMonitorTxUnion.robotMonitorTxData.dataNum = robotMonitorRxUnion.robotMonitorRxData.dataNum;
  robotMonitorTxUnion.robotMonitorTxData.carriageReturn = '\n';

  /* Send the error message */
  HAL_UART_Transmit_DMA(&RM_UART, robotMonitorTxUnion.robotMonitorTxBuf, 9);
}

/**
 * @brief  Subscribes the variable in flash memory of the given address.
 * @param  None
 * @retval None
 */
void subscribeFlash(void)
{
  /* Clear the data buffer to be sent */
  memset((uint8_t *) &robotMonitorTxUnion.robotMonitorTxBuf, 0, sizeof(robotMonitorTxUnion.robotMonitorTxBuf));
  memset((uint8_t *) &txDataBuf, 0, sizeof(txDataBuf));
  txDataLen = 0;

  /* Prepare the data to send */
  robotMonitorTxUnion.robotMonitorTxData.board = robotMonitorRxUnion.robotMonitorRxData.board;
  robotMonitorTxUnion.robotMonitorTxData.act = RM_ACT_SUBSCRIBERETURN;
  robotMonitorTxUnion.robotMonitorTxData.carriageReturn = '\n';

  uint16_t dataNum;
  uint32_t addr;
  for (int i = 0; i < MAX_ADDR_NUM; i++)
  {
    if (listAddr[i].dataNum != 0)
    {
      addr = listAddr[i].addr;
      robotMonitorTxUnion.robotMonitorTxData.addr = listAddr[i].addr;
      /* Reads the variable in flash memory */
      for (dataNum = 0; dataNum < listAddr[i].dataNum; dataNum++)
      {
        *(robotMonitorTxUnion.robotMonitorTxBuf + 7 + dataNum) = *(__IO uint8_t*) addr++;
      }
      robotMonitorTxUnion.robotMonitorTxData.dataNum = dataNum;
      robotMonitorTxUnion.robotMonitorTxData.tick = osKernelGetTickCount();
      memcpy((uint8_t *)(txDataBuf + txDataLen), (uint8_t *)robotMonitorTxUnion.robotMonitorTxBuf, 20);
      txDataLen += 20;
    }
  }
  /* Send return data */
  HAL_UART_Transmit_DMA(&RM_UART, (uint8_t *)txDataBuf, txDataLen);
}

/**
 * @brief  Writes the given data buffer to the flash of the given address.
 * @param  None
 * @retval None
 */
void writeFlash(void)
{
  /* Clear the data buffer to be sent */
  memset((uint8_t *) &robotMonitorTxUnion.robotMonitorTxBuf, 0, sizeof(robotMonitorTxUnion.robotMonitorTxBuf));

  /* Prepare the data to send */
  robotMonitorTxUnion.robotMonitorTxData.board = robotMonitorRxUnion.robotMonitorRxData.board;
  robotMonitorTxUnion.robotMonitorTxData.addr = robotMonitorRxUnion.robotMonitorRxData.addr;
  robotMonitorTxUnion.robotMonitorTxData.dataNum = robotMonitorRxUnion.robotMonitorRxData.dataNum;
  robotMonitorTxUnion.robotMonitorTxData.carriageReturn = '\n';

  /* Write data buffer */
  uint32_t TypeProgram;
  switch (robotMonitorRxUnion.robotMonitorRxData.dataNum)
  {
    case 1:
      TypeProgram = FLASH_TYPEPROGRAM_BYTE;
      break;
    case 2:
      TypeProgram = FLASH_TYPEPROGRAM_HALFWORD;
      break;
    case 4:
      TypeProgram = FLASH_TYPEPROGRAM_WORD;
      break;
    case 8:
      TypeProgram = FLASH_TYPEPROGRAM_DOUBLEWORD;
      break;
    default:
      returnError(RM_ERROR_NOSUCHDATANUM);
      return;
  }
  if (HAL_FLASH_Unlock() == HAL_OK)
    if (HAL_FLASH_Program(TypeProgram, robotMonitorRxUnion.robotMonitorRxData.addr, robotMonitorRxUnion.robotMonitorRxData.dataBuf) == HAL_OK)
      if (HAL_FLASH_Lock() == HAL_OK)
        robotMonitorTxUnion.robotMonitorTxData.act = RM_ACT_WRITERETURN;

  /* Send return data */
  HAL_UART_Transmit_DMA(&RM_UART, robotMonitorTxUnion.robotMonitorTxBuf, 8);
}

/**
 * @brief  Function implementing the robotMonitorTask thread
 * @param  argument: Not used
 * @retval None
 */
void tskRobotMonitor(void *argument)
{
  uint32_t tick;
  tick = osKernelGetTickCount();
  HAL_UART_Receive_DMA(&RM_UART, robotMonitorRxUnion.robotMonitorRxBuf, sizeof(robotMonitorRxUnion.robotMonitorRxBuf));
  __HAL_UART_ENABLE_IT(&RM_UART, UART_IT_IDLE);
  /* Infinite loop */
  for (;;)
  {
    /* Delay 1000 ticks periodically */
    tick += 10U;
    osDelayUntil(tick);

    subscribeFlash();
    /* Check if data is received */
    if (isGetRobotMonitor)
    {

      /* Check if it is a valid board ID */
      if (robotMonitorRxUnion.robotMonitorRxData.board == RM_BOARD_1)
      {

        /* Check if it is a valid action information and execute */
        if (robotMonitorRxUnion.robotMonitorRxData.act == RM_ACT_SUBSCRIBE)
        {
          if (addrRegister() == -1)
          {
            returnError(RM_ERROR_FULLADDR);
          }
        }
        else if (robotMonitorRxUnion.robotMonitorRxData.act == RM_ACT_WRITE)
        {
          writeFlash();
        }
        else if (robotMonitorRxUnion.robotMonitorRxData.act == RM_ACT_UNSUBSCRIBE)
        {
          if (addrUnregister() == -1)
          {
            returnError(RM_ERROR_NOSUCHADDRREG);
          }
        }
        else
        {
          returnError(RM_ERROR_NOSUCHACT);
        }

      }
      else if (robotMonitorRxUnion.robotMonitorRxData.board == RM_BOARD_2 || robotMonitorRxUnion.robotMonitorRxData.board == RM_BOARD_3)
      {
        //TODO
      }
      else
      {
        returnError(RM_ERROR_NOSUCHBOARD);
      }
      isGetRobotMonitor = 0;
    }
  }
}
