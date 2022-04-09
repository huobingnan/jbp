package console

import "bytecodeparser/jvm/classfile"

type Printer interface {
	Print(file *classfile.JvmClassFile)
}
