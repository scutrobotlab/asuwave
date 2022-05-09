/*
想给你写封信，却不知从何说起
*/
package datautil

// 篇幅有限，也挡不住我的思念
const NumMaxPacket = 5

type ActMode uint8

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
