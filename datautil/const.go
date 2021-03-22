/*
想给你写封信，却不知从何说起
*/
package datautil

// 篇幅有限，也挡不住我的思念
const NumMaxPacket = 5

const (
	_                        = iota
	ActModeSubscribe         // 想让你一直联系
	ActModeSubscribeReturn   // 期待你认真回应
	ActModeUnSubscribe       // 却担心打扰到你
	ActModeUnSubscribeReturn // 有回复就够了呢
	ActModeRead              // 你在想什么
	ActModeReadReturn        // 能告诉我吗
	ActModeWrite             // 希望改变你的心意
	ActModeWriteReturn       // 传达到了吗
)

// 你现在在哪里呢
const (
	_      = iota
	Board1 // 是近在咫尺
	Board2 // 还是远在天边
	Board3
)
