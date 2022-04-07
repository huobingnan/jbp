package classfile

import (
	"bytecodeparser/jvm/classfile/reader"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestNewJvmClassFile(t *testing.T) {

	// NOTICE: 测试前，请将data目录在本机的绝对路径拷贝在这里
	testingDataDir := "/Users/huobingnan/Code/Golang/ByteCodeParser/data"

	t.Run("BiLock.class", func(t *testing.T) {
		byteCode, err := ioutil.ReadFile(filepath.Join(testingDataDir, "BiLock.class"))
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
		if jvmClassFile.fields != nil {
			for _, each := range jvmClassFile.fields {
				t.Log("Filed:", &each)
			}
		}
		for _, method := range jvmClassFile.Methods() {
			t.Log("Method:", &method)
		}
		jvmClassFile.Print(new(SimpleConsolePrinter))
	})

	t.Run("InterfaceDemo.class", func(t *testing.T) {
		byteCode, _ := ioutil.ReadFile(filepath.Join(testingDataDir, "InterfaceDemo.class"))
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
		t.Logf("Interface index set: %v", jvmClassFile.InterfaceIndexSet())
		if jvmClassFile.Fields() != nil {
			for _, each := range jvmClassFile.Fields() {
				t.Log("Field:", &each)
			}
		}
		for _, method := range jvmClassFile.Methods() {
			t.Log("Method:", &method)
		}
	})

	t.Run("GroovyClosureCurrying.class", func(t *testing.T) {
		byteCode, _ := ioutil.ReadFile(filepath.Join(testingDataDir, "GroovyClosureCurrying.class"))
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
		t.Logf("Interface index set: %v", jvmClassFile.InterfaceIndexSet())
		if jvmClassFile.Fields() != nil {
			for _, each := range jvmClassFile.Fields() {
				t.Log("Field:", &each)
			}
		}
	})

	t.Run("TurnLock.class", func(t *testing.T) {
		var jvmClassFile *JvmClassFile
		byteCode, _ := ioutil.ReadFile(filepath.Join(testingDataDir, "TurnLock.class"))
		r := reader.NewByteCodeReader(byteCode)
		jvmClassFile = NewJvmClassFile(r)
		t.Logf("Magic number: %X", jvmClassFile.magicNumber)
		t.Log("Minor version: ", jvmClassFile.minorVersion)
		t.Log("Major version: ", jvmClassFile.majorVersion)
		for idx, each := range jvmClassFile.cp {
			t.Logf("index: %2d, %v", idx, each)
		}
		t.Logf("Access flags: %x", jvmClassFile.accessFlags)
		t.Logf("This class: %v", jvmClassFile.ConstantPool()[jvmClassFile.thisClass])
		t.Logf("Super class: %v", jvmClassFile.ConstantPool()[jvmClassFile.superClass])
		t.Logf("Interface index set: %v", jvmClassFile.InterfaceIndexSet())
		if jvmClassFile.Fields() != nil {
			for _, each := range jvmClassFile.Fields() {
				t.Log("Field:", &each)
			}
		}
		for _, method := range jvmClassFile.Methods() {
			t.Log("Method:", &method)
		}
	})

}
