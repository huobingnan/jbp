package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestNewConstantPool(t *testing.T) {
	t.Run("BiLock.class", func(t *testing.T) {
		byteCode, _ := ioutil.ReadFile("/Users/huobingnan/Code/Golang/ByteCodeParser/data/BiLock.class")
		r := reader.NewByteCodeReader(byteCode)
		// 读取魔数
		magic, _ := r.ReadU4()
		r.ReadU2() // 读取主版本号
		r.ReadU2() // 读取次版本号
		fmt.Printf("%X\n", magic)
		constantPool := NewConstantPool(r)
		for idx, each := range constantPool {
			fmt.Printf("index: %2d, %v\n", idx, each)
		}
	})

	t.Run("TurnLock.class", func(t *testing.T) {
		byteCode, _ := ioutil.ReadFile("/Users/huobingnan/Code/Golang/ByteCodeParser/data/TurnLock.class")
		r := reader.NewByteCodeReader(byteCode)
		r.ReadU4() // 读取魔数
		r.ReadU2() // 读取主版本号
		r.ReadU2() // 读取次版本号
		cp := NewConstantPool(r)
		for idx, each := range cp {
			fmt.Printf("index: %2d, %v\n", idx, each)
		}
	})
}
