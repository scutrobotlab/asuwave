/// UTF-8 Encode
/**
 * @file SerialLineIP.hpp
 * @author BigeYoung (SCUT.BigeYoung@gmail.com)
 * @brief SerialLineIP 是一种简单的数据链路层串口协议，
 *  提供了封装成帧和透明传输的功能。
 * @warning STANDARD C++03 REQUIRED! 需要C++03以上标准支持！
 *  - C++03起，vector 元素相继存储，因此，您能用指向元素的常规指针
 *  访问元素。例如，对于返回值data，您可以使用&data[0]获取数组首地址。
 *  [C++标准:vector](https://zh.cppreference.com/w/cpp/container/vector)
 *  - 依照 ISO C++ 标准建议，本程序的函数直接返回 vector 容器。
 *  在支持具名返回值优化(NRVO)的编译器上可获得最佳性能。
 *  [C++标准:复制消除](https://zh.cppreference.com/w/cpp/language/copy_elision)
 * @see [RFC-1055: SLIP 协议文档](https://tools.ietf.org/html/rfc1055)
 * @version 0.1
 * @date 2018-12-24
 * 
 * @copyright Copyright 华工机器人实验室(c) 2018
 * 
 */
#ifndef SERIAL_LINE_IP_H
#define SERIAL_LINE_IP_H
#include <stdio.h>
#include <stdint.h> /* uint8_t */
#include <assert.h> /* assert */
#include <vector>   /* vector */
namespace SerialLineIP
{
/* SLIP special character codes */
const uint8_t END = 0xC0;     /* indicates end of packet */
const uint8_t ESC = 0xDB;     /* indicates byte stuffing */
const uint8_t ESC_END = 0xDC; /* ESC ESC_END means END data byte */
const uint8_t ESC_ESC = 0xDD; /* ESC ESC_ESC means ESC data byte */

/**
 * @brief Serial Line IP PACK
 * @param p_PDU 起始位置指针
 * @param PDU_len PDU 字节长度
 * @return std::vector<uint8_t> SDU
 * @note Service Data Unit (SDU) 指本层封包后产生的数据单元
 *       Protocol Data Unit (PDU) 指上层协议数据单元
 */
std::vector<uint8_t> Pack(const void *const p_PDU, int PDU_len);

/**
 * @brief Serial Line IP UNPACK
 * @param p_SDU 起始位置指针
 * @param SDU_len SDU 字节长度
 * @return std::vector<uint8_t> PDU
 * @note Service Data Unit (SDU) 指本层解包前的数据单元
 *       Protocol Data Unit (PDU) 指上层协议数据单元
 */
std::vector<uint8_t> Unpack(const void *const p_SDU, int SDU_len);
} // namespace SerialLineIP
#endif // SERIAL_LINE_IP_H