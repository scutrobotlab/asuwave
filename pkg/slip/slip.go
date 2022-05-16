/**
 * SerialLineIP 是一种简单的数据链路层串口协议，提供了封装成帧和透明传输的功能。
 * 请参阅：[RFC-1055: SLIP 协议文档](https://tools.ietf.org/html/rfc1055)
**/

package slip

import (
	"errors"

	"github.com/golang/glog"
)

/* SLIP special character codes */
const (
	END     byte = 0xC0 // indicates end of packet
	ESC     byte = 0xDB // indicates byte stuffing
	ESC_END byte = 0xDC // ESC ESC_END means END data byte
	ESC_ESC byte = 0xDD // ESC ESC_ESC means ESC data byte
)

/**
 * Serial Line IP PACK
 * Service Data Unit (SDU) 指本层封包后产生的数据单元
 * Protocol Data Unit (PDU) 指上层协议数据单元
 */
func Pack(PDU []byte) (SDU []byte) {
	SDU_len := len(PDU) + 2
	for _, p := range PDU {
		if p == END || p == ESC {
			SDU_len++
		}
	}
	SDU = make([]byte, 0, SDU_len)

	SDU = append(SDU, END)
	for _, p := range PDU {
		switch p {
		case END:
			SDU = append(SDU, ESC, ESC_END)
		case ESC:
			SDU = append(SDU, ESC, ESC_ESC)
		default:
			SDU = append(SDU, p)
		}
	}
	SDU = append(SDU, END)
	if len(SDU) == SDU_len {
		glog.Errorf("SLIP: length %d != expected length %d", len(SDU), SDU_len)
	}
	return SDU
}

/**
 * Serial Line IP UNPACK
 * Service Data Unit (SDU) 指本层封包后产生的数据单元
 * Protocol Data Unit (PDU) 指上层协议数据单元
 */
func Unpack(SDU []byte) (PDU []byte, err error) {
	PDU = make([]byte, 0, len(SDU)-2)
	i := 0
	for i = 0; i < len(SDU); i++ {
		if SDU[i] == END {
			break
		}
	}
	for i++; i < len(SDU); i++ {
		switch SDU[i] {
		case END:
			return PDU, nil
		case ESC:
			i++
			switch SDU[i] {
			case ESC_END:
				PDU = append(PDU, END)
			case ESC_ESC:
				PDU = append(PDU, ESC)
			default:
				return PDU, errors.New("unknown byte after ESC")
			}
		default:
			PDU = append(PDU, SDU[i])
		}
	}
	return []byte{}, errors.New("END unfound")
}
