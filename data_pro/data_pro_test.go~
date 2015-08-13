package pack_pro

import (
	"fmt"
	"testing"
)

func Test_Pack(t *testing.T) {
	fmt.Println([]byte("hello world!"))
	re, ok := Pack([]byte("hello world!"), 12)
	if ok {
		fmt.Println("1:", re)
	}
	a, err := UnPack(re)
	fmt.Println(err)
	fmt.Println("2:", a)
	if err == nil {
		fmt.Println("3:", string(a[0].([]byte)))
	}
}
