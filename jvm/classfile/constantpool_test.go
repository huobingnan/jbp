package classfile

import (
	"bytecodeparser/jvm/reader"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestNewConstantPool(t *testing.T) {

	// NOTICE: 测试前，请将data目录在本机的绝对路径拷贝在这里
	testingDataDir := "/Users/huobingnan/Code/Golang/ByteCodeParser/data"

	t.Run("BiLock.class", func(t *testing.T) {
		byteCode, _ := ioutil.ReadFile(filepath.Join(testingDataDir, "BiLock.class"))
		r := reader.NewByteCodeReader(byteCode)
		// 读取魔数
		magic, _ := r.ReadU4()
		r.ReadU2() // 读取主版本号
		r.ReadU2() // 读取次版本号
		fmt.Printf("%X\n", magic)
		constantPool := NewConstantPool(r)
		for idx, each := range constantPool {
			fmt.Printf("Index: %2d, %v\n", idx, each)
		}
	})

	t.Run("TurnLock.class", func(t *testing.T) {
		byteCode, _ := ioutil.ReadFile(filepath.Join(testingDataDir, "TurnLock.class"))
		r := reader.NewByteCodeReader(byteCode)
		r.ReadU4() // 读取魔数
		r.ReadU2() // 读取主版本号
		r.ReadU2() // 读取次版本号
		cp := NewConstantPool(r)
		for idx, each := range cp {
			fmt.Printf("Index: %2d, %v\n", idx, each)
		}
	})
}
