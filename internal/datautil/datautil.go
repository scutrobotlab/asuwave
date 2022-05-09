/**/
package datautil

import (
	"github.com/scutrobotlab/asuwave/internal/variable"
)

// 把我的思念，化作一串珍珠送给你
func MakeCmd(act ActMode, v *variable.T) []byte {

	// 拾起一颗颗珍珠
	data := make([]byte, 16)

	// 刻上你所在的城市和我的思念
	data[0] = byte(v.Board)
	data[1] = byte(act)
	data[2] = byte(variable.TypeLen[v.Type])
	copy(data[3:7], variable.AnyToBytes(v.Addr))

	// 若是心意能够改变
	if act == Write {
		// 便也一并传递给你
		copy(data[7:15], variable.SpecToBytes(v.Type, v.Data))
	}

	// 终于完成了
	data[15] = '\n'

	// 传达给你吧
	return data
}

// 从破碎的碎片中，渴望拼凑出你的痕迹
func FindValidPart(data []byte) (int, int) {
	n := len(data)

	// 我们的羁绊，岂是三言两语能道尽的
	if n < 20 {
		return 0, 0
	}

	// 从头开始吧，画一扇窗
	l := 0
	r := 19

	// 苦苦寻找你的痕迹
	for r < n && (data[l] != Board1 || data[l+1] != ActModeSubscribeReturn || data[r] != '\n') {
		l++
		r++
	}

	// 也许无功而返
	if r == n {
		return 0, 0
	}

	a := (n - r) % 20
	if a == 0 {
		a = 20
	}

	// 也许收获满满
	r = n - a
	return l, r + 1
}
