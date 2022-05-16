/**/
package variable

// 把我的思念，化作一串珍珠送给你
func MakeWriteCmd(v T) []byte {
	// 拾起一颗颗珍珠
	data := make([]byte, 16)

	// 刻上你所在的城市和我的思念
	data[0] = byte(v.Board)
	data[1] = byte(Write)
	data[2] = byte(TypeLen[v.Type])
	copy(data[3:7], AnyToBytes(v.Addr))

	copy(data[7:15], SpecToBytes(v.Type, v.Data))

	// 终于完成了
	data[15] = '\n'

	// 传达给你吧
	return data
}

// 把我的思念，化作一串珍珠送给你
func MakeCmd(act ActMode, v CmdT) []byte {

	// 拾起一颗颗珍珠
	data := make([]byte, 16)

	// 刻上你所在的城市和我的思念
	data[0] = byte(v.Board)
	data[1] = byte(act)
	data[2] = byte(v.Length)
	copy(data[3:7], AnyToBytes(v.Addr))

	// 终于完成了
	data[15] = '\n'

	// 传达给你吧
	return data
}
