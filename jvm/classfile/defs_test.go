package classfile

import (
	"bytecodeparser/jvm/classfile/reader"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func _basicallyPrintClassFile(file string, t *testing.T) {
	byteCode, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	r := reader.NewByteCodeReader(byteCode)
	jvmClassFile := NewJvmClassFile(r)
	t.Logf("Magic number: %X", jvmClassFile.magicNumber)
	t.Log("Minor version: ", jvmClassFile.minorVersion)
	t.Log("Major version: ", jvmClassFile.majorVersion)
	for idx, each := range jvmClassFile.cp {
		t.Logf("index: %2d, %v", idx, each)
	}
	t.Logf("Access flags: %x", jvmClassFile.accessFlags)
	t.Logf("This class: %v", jvmClassFile.ConstantPool()[jvmClassFile.thisClass])
	t.Logf("Super class: %v", jvmClassFile.ConstantPool()[jvmClassFile.superClass])
	for _, field := range jvmClassFile.Fields() {
		t.Log("Field:", &field)
	}
	for _, method := range jvmClassFile.Methods() {
		t.Log("Method:", &method)
	}
	for _, attr := range jvmClassFile.Attributes() {
		t.Log(attr.Name())
	}
}

func TestNewJvmClassFile(t *testing.T) {

	// NOTICE: 测试前，请将data目录在本机的绝对路径拷贝在这里
	testingDataDir := "/Users/huobingnan/Code/Golang/ByteCodeParser/data"

	t.Run("BiLock.class", func(t *testing.T) {
		_basicallyPrintClassFile(filepath.Join(testingDataDir, "BiLock.class"), t)
	})

	t.Run("InterfaceDemo.class", func(t *testing.T) {
		_basicallyPrintClassFile(filepath.Join(testingDataDir, "InterfaceDemo.class"), t)
	})

	t.Run("GroovyClosureCurrying.class", func(t *testing.T) {
		_basicallyPrintClassFile(filepath.Join(testingDataDir, "GroovyClosureCurrying.class"), t)
	})

	t.Run("TurnLock.class", func(t *testing.T) {
		_basicallyPrintClassFile(filepath.Join(testingDataDir, "TurnLock.class"), t)
	})

	t.Run("Exceptions.class", func(t *testing.T) {
		_basicallyPrintClassFile(filepath.Join(testingDataDir, "Exceptions.class"), t)
	})

	t.Run("InnerClass.class", func(t *testing.T) {
		_basicallyPrintClassFile(filepath.Join(testingDataDir, "InnerClass.class"), t)
	})

}
