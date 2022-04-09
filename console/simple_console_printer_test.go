package console

import (
	"bytecodeparser/jvm/classfile"
	"bytecodeparser/jvm/classfile/reader"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestSimpleConsolePrinter_Print(t *testing.T) {
	testingDataDir := "/Users/huobingnan/Code/Golang/ByteCodeParser/data"
	t.Run("BiLock.class", func(t *testing.T) {
		b, err := ioutil.ReadFile(filepath.Join(testingDataDir, "BiLock.class"))
		if err != nil {
			t.Fatal(err)
		}
		r := reader.NewByteCodeReader(b)
		cfile := classfile.NewJvmClassFile(r)
		printer := NewSimpleConsolePrinter(cfile)
		printer.Print()
	})
}
