package classfile

import (
	"bytecodeparser/jvm/classfile/reader"
	"io/ioutil"
	"testing"
)

func TestNewJvmClassFile(t *testing.T) {
	t.Run("BiLock.class", func(t *testing.T) {
		byteCode, _ := ioutil.ReadFile("/Users/huobingnan/Code/Golang/ByteCodeParser/data/BiLock.class")
		r := reader.NewByteCodeReader(byteCode)
		jvmClassFile := NewJvmClassFile(r)
		t.Log("Magic number: ", jvmClassFile.magicNumber)
		t.Log("Minor version: ", jvmClassFile.minorVersion)
		t.Log("Major version: ", jvmClassFile.majorVersion)
		for idx, each := range jvmClassFile.cp {
			t.Logf("index: %2d, %v", idx, each)
		}
		t.Logf("Access flags: %x", jvmClassFile.accessFlags)
		t.Logf("This class: %v", jvmClassFile.ConstantPool()[jvmClassFile.thisClass])
		t.Logf("Super class: %v", jvmClassFile.ConstantPool()[jvmClassFile.superClass])
	})

	t.Run("InterfaceDemo.class", func(t *testing.T) {
		byteCode, _ := ioutil.ReadFile("/Users/huobingnan/Code/Golang/ByteCodeParser/data/InterfaceDemo.class")
		r := reader.NewByteCodeReader(byteCode)
		jvmClassFile := NewJvmClassFile(r)
		t.Log("Magic number: ", jvmClassFile.magicNumber)
		t.Log("Minor version: ", jvmClassFile.minorVersion)
		t.Log("Major version: ", jvmClassFile.majorVersion)
		for idx, each := range jvmClassFile.cp {
			t.Logf("index: %2d, %v", idx, each)
		}
		t.Logf("Access flags: %x", jvmClassFile.accessFlags)
		t.Logf("This class: %v", jvmClassFile.ConstantPool()[jvmClassFile.thisClass])
		t.Logf("Super class: %v", jvmClassFile.ConstantPool()[jvmClassFile.superClass])
		t.Logf("Interface index set: %v", jvmClassFile.InterfaceIndexSet())
	})

	t.Run("GroovyClosureCurrying.class", func(t *testing.T) {
		byteCode, _ := ioutil.ReadFile("/Users/huobingnan/Code/Golang/ByteCodeParser/data/GroovyClosureCurrying.class")
		r := reader.NewByteCodeReader(byteCode)
		jvmClassFile := NewJvmClassFile(r)
		t.Log("Magic number: ", jvmClassFile.magicNumber)
		t.Log("Minor version: ", jvmClassFile.minorVersion)
		t.Log("Major version: ", jvmClassFile.majorVersion)
		for idx, each := range jvmClassFile.cp {
			t.Logf("index: %2d, %v", idx, each)
		}
		t.Logf("Access flags: %x", jvmClassFile.accessFlags)
		t.Logf("This class: %v", jvmClassFile.ConstantPool()[jvmClassFile.thisClass])
		t.Logf("Super class: %v", jvmClassFile.ConstantPool()[jvmClassFile.superClass])
		t.Logf("Interface index set: %v", jvmClassFile.InterfaceIndexSet())
	})

}
