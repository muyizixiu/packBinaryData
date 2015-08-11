package pack

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

var classTable map[string][]byte = map[string][]byte{
	"number": {1},
	"string": {2},
	"json":   {3},
	"bytes":  {4},
}

const (
	headerLen = 4
	classLen  = 1
	tailLen   = 4
)

func Pack(data ...interface{}) ([]byte, error) {
	data_re := make([]byte, 100)
	for _, v := range data {
		data_tmp, err := pack(v)
		if err != nil {
			return nil, err
		}
		data_re = append(data_re, data_tmp...)
	}
	var i int
	for {
		if data_re[i] != 0 {
			data_re = data_re[i:]
			break
		}
		i++
	}
	return data_re, nil
}
func pack(data interface{}) ([]byte, error) {
	var data_re []byte
	switch data.(type) {
	case int:
		data_tmp := convertInt(data.(int))
		test(data_tmp)
		data_re = class(data_tmp, "number")
	case string:
		data_re = class([]byte(data.(string)), "string")
	case []byte:
		data_re = class(data.([]byte), "bytes")
	default:
		data_re = nil
		return nil, errors.New("error: don't support this data type")
	}
	if data != nil {
		if data_re == nil {
			return nil, errors.New("internal error!")
		}
	}
	data_re = header(data_re)
	data_re = tail(data_re)
	return data_re, nil
}
func class(data []byte, c string) []byte {
	classInfo, _ := classTable[c]
	for len(classInfo) != classLen {
		if len(classInfo) > classLen {
			errorLog(errors.New("there are value out of range in classTable!"))
			return nil
		}
		classInfo = append([]byte{0}, classInfo...)
	}
	return append(classInfo, data...)

}
func header(data []byte) []byte {
	length := int32(len(data))
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, length)
	if err != nil {
		errorLog(err)
		return nil
	}
	return append(buf.Bytes(), data...)
}
func tail(data []byte) []byte {
	return data
}
func errorLog(err error) {
	fmt.Println(err.Error())
}
func test(v interface{}) {
	fmt.Println("\n", v, "\n")
}
func convertInt(n int) []byte {
	buf := new(bytes.Buffer)
	a := int64(n)
	err := binary.Write(buf, binary.LittleEndian, a)
	if err != nil {
		errorLog(err)
		return nil
	}
	return rmZero(buf.Bytes(), false)
}
func rmZero(k []byte, flag bool) []byte {
	length := len(k) - 1
	if flag {
		for i, v := range k {
			if v != 0 {
				return k[i:]
			}
		}
	} else {
		for i, _ := range k {
			if k[length-i] != 0 {
				return k[:(length - i + 1)]
			}
		}
	}
	return nil
}
