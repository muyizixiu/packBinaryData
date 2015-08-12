package pack

import (
	"fmt"
	"testing"
)

func _Pack(v ...interface{}) {
	data, err := Pack(v...)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(data)
}
func Test_pack(t *testing.T) {
	_Pack(3)
	_Pack("hello world!")
	buf := []byte{104, 101, 108, 111}
	_Pack(buf)
	_Pack(3, "hello world!", buf)
}
