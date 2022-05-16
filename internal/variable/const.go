/*
想给你写封信，却不知从何说起
*/
package variable

// 篇幅有限，也挡不住我的思念
const NumMaxPacket = 5

type ActMode uint8

var TypeLen = map[string]int{
	"uint8_t":  1,
	"uint16_t": 2,
	"uint32_t": 4,
	"uint64_t": 8,
	"int8_t":   1,
	"int16_t":  2,
	"int32_t":  4,
	"int64_t":  8,
	"int":      4,
	"float":    4,
	"double":   8,
}

const (
	_                 ActMode = iota
	Subscribe                 // 想让你一直联系
	SubscribeReturn           // 期待你认真回应
	Unsubscribe               // 却担心打扰到你
	UnsubscribeReturn         // 有回复就够了呢
	Read                      // 你在想什么
	ReadReturn                // 能告诉我吗
	Write                     // 希望改变你的心意
	WriteReturn               // 传达到了吗
)

// 你现在在哪里呢
const (
	_      = iota
	Board1 // 是近在咫尺
	Board2 // 还是远在天边
	Board3
)
