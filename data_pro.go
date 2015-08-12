package pack_pro

import (
	"errors"
)

var (
	version     int8 = 1
	sub_version int8 = 1
	headLen          = 9
)

func Pack(v ...interface{}) ([]byte, bool) {
}
func pack(v []byte, cmd int16) ([]byte, bool) {
	v_tmp := encryEnd(v)
	v_tmp = WriteCmd(cmd, v_tmp)
	v_tmp = WriteInt8(sub_version, v_tmp)
	v_tmp = WriteInt8(version, v_tmp)
	v_tmp = WriteChar("Y", v_tmp)
	v_tmp = WriteChar("B", v_tmp)
	v_tmp = WriteLen(v_tmp)
}
func checkErr(err error) bool {
	if err != nil {
		fmt.Println("\n", err.Error(), "\n")
		return true
	}
	return false
}
func WriteInt8(i int8, v []byte) []byte {
	b := make([]byte, 1)
	b[0] = i
	b = append(b, v)
	return b
}
func WriteChar(s string, v []byte) []byte {
	b := []byte(s)
	return append(b[:1], v...)
}
func encryEnd(v []byte) []byte {
}
func WriteCmd(cmd int16, v []byte) []byte {
	return WriteInt8(cmd/256, WriteInt8(cmd%256, v))
}
func encrypt(v []byte) []byte {
}
func WriteLen(v []byte) []byte {
	length := int16(len(v) - 7)
	return WriteCmd(length, v)
}
func UnPack(data []byte) ([]interface{}, error) {
}
func unpack(data []byte) ([]byte, int16, error) {
	BY := []byte{'B', 'Y', version, sub_version}
	for i, v := range BY {
		if data[i] != v {
			return nil, 0, errors.New("flags are not right!")
		}
	}
	var cmd int16
	cmd = int16(data[4])*256 + int16(data[5])
}
