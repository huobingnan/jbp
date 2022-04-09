package console

import (
	"bytecodeparser/jvm/classfile"
	"bytecodeparser/jvm/classfile/constantpool"
	"fmt"
)

type SimpleConsolePrinter struct {
}

func _printConstantPool(cfile *classfile.JvmClassFile) {
	fmt.Printf("Constant Pool:\n")
	pool := cfile.ConstantPool()
	var index1, index2 uint16
	for idx, cp := range pool {
		if cp == nil {
			continue
		}
		switch cp.Tag() {
		case constantpool.ConstantUtf8Info:
			fmt.Printf("\t#%d = %s // utf8\n", idx, cp.Value().(string))
		case constantpool.ConstantIntegerInfo:
			fmt.Printf("\t#%d = %d // integer\n", idx, cp.Value().(int))
		case constantpool.ConstantLongInfo:
			fmt.Printf("\t#%d = %d // long\n", idx, cp.Value().(int64))
		case constantpool.ConstantFloatInfo:
			fmt.Printf("\t#%d = %f // float\n", idx, cp.Value().(float32))
		case constantpool.ConstantDoubleInfo:
			fmt.Printf("\t#%d = %f // double\n", idx, cp.Value().(float64))
		case constantpool.ConstantStringInfo:
			index1 = cp.Value().(uint16)
			fmt.Printf("\t#%d = %d // string; %s\n", idx, index1, pool[index1].Value().(string))
		case constantpool.ConstantClassInfo:
			index1 = cp.Value().(uint16)
			fmt.Printf("\t#%d = @%d // class; %s\n", idx, index1, pool[index1].Value().(string))
		case constantpool.ConstantMethodRefInfo:
			index1 = cp.Value().([]uint16)[0] // class info
			index2 = cp.Value().([]uint16)[1] // name and type
			fmt.Printf("\t#%d = @%d, @%d // methodref; ", idx, index1, index2)
			index1 = pool[index1].Value().(uint16) // class name utf8 index
			className := pool[index1].Value().(string)
			index1 = pool[index2].Value().([]uint16)[0]
			methodName := pool[index1].Value().(string)
			index1 = pool[index2].Value().([]uint16)[1]
			descriptor := pool[index1].Value().(string)
			fmt.Printf("%s.%s%s\n", className, methodName, descriptor)

		case constantpool.ConstantFieldRefInfo:
			index1 = cp.Value().([]uint16)[0] // class info
			index2 = cp.Value().([]uint16)[1] // name and type
			fmt.Printf("\t#%d = @%d, %d // fieldref; ", idx, index1, index2)
			index1 = pool[index1].Value().(uint16) // class name utf8 index
			className := pool[index1].Value().(string)
			index1 = pool[index2].Value().([]uint16)[0]
			methodName := pool[index1].Value().(string)
			index1 = pool[index2].Value().([]uint16)[1]
			descriptor := pool[index1].Value().(string)
			fmt.Printf("%s.%s%s\n", className, methodName, descriptor)
		}
	}
}

func (s *SimpleConsolePrinter) Print(cfile *classfile.JvmClassFile) {
	if cfile == nil {
		return
	}
	_printConstantPool(cfile)
}
