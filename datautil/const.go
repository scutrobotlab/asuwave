package datautil

const MAX_PACKET_NUM = 5

const (
	_ = iota
	ActModeSubscribe
	ActModeSubscribeReturn
	ActModeUnSubscribe
	ActModeUnSubscribeReturn
	ActModeRead
	ActModeReadReturn
	ActModeWrite
	ActModeWriteReturn
)

const (
	_ = iota
	Board1
	Board2
	Board3
)
