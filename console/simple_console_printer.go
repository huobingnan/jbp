package console

import (
	"bytecodeparser/jvm/classfile"
	"bytecodeparser/jvm/classfile/constantpool"
	"fmt"
)

type SimpleConsolePrinter struct {
	classFile *classfile.JvmClassFile
}

func (s *SimpleConsolePrinter) _parseClassName(cidx uint16) string {
	cp := s.classFile.ConstantPool()
	classNameIndex := cp[cidx].Value().(uint16)
	return cp[classNameIndex].Value().(string)
}

func (s *SimpleConsolePrinter) _printConstantPool() {
	cp := s.classFile.ConstantPool()
	fmt.Printf("Constant Pool:\n")
	for idx, constant := range cp {
		if constant == nil {
			continue
		}
		switch constant.Tag() {
		case constantpool.ConstantIntegerInfo:
			fmt.Printf("\t#%d = %d // Integer;\n", idx, constant.Value().(int))
		case constantpool.ConstantLongInfo:
			fmt.Printf("\t#%d = %d // Long;\n", idx, constant.Value().(int64))
		case constantpool.ConstantFloatInfo:
			fmt.Printf("\t#%d = %f // Float;\n", idx, constant.Value().(float32))
		case constantpool.ConstantDoubleInfo:
			fmt.Printf("\t#%d = %f // Double;\n", idx, constant.Value().(float64))
		case constantpool.ConstantStringInfo:
			s := constant.(*constantpool.StringConstant)
			fmt.Printf("\t#%d = @%d // String; %s\n", idx, s.Index, cp[s.Index].Value().(string))
		case constantpool.ConstantUtf8Info:
			fmt.Printf("\t#%d = %s // Utf8;\n", idx, constant.Value().(string))
		case constantpool.ConstantNameAndTypeInfo:
			info := constant.(*constantpool.NameAndTypeConstant)
			fmt.Printf("\t#%d = @%d, @%d // NameAndType; %s %s\n", idx,
				info.NameIndex,
				info.DescriptorIndex,
				cp[info.NameIndex].Value().(string),
				cp[info.DescriptorIndex].Value().(string))
		case constantpool.ConstantClassInfo:
			info := constant.(*constantpool.ClassConstant)
			fmt.Printf("\t#%d = @%d // Class; %s\n", idx, info.Index, cp[info.Index].Value().(string))
		case constantpool.ConstantMethodRefInfo:
			info := constant.(*constantpool.MethodRefConstant)
			desc := cp[info.NameAndTypeInfoIndex].(*constantpool.NameAndTypeConstant)
			fmt.Printf("\t#%d = @%d, @%d // Methodref; %s.%s%s\n", idx, info.ClassInfoIndex,
				info.NameAndTypeInfoIndex, s._parseClassName(info.ClassInfoIndex),
				cp[desc.NameIndex].Value().(string), cp[desc.DescriptorIndex].Value().(string))
		case constantpool.ConstantFieldRefInfo:
			info := constant.(*constantpool.FieldRefConstant)
			desc := cp[info.NameAndTypeInfoIndex].(*constantpool.NameAndTypeConstant)
			fmt.Printf("\t#%d = @%d, @%d // Fieldref; %s.%s %s\n", idx, info.ClassInfoIndex,
				info.NameAndTypeInfoIndex, s._parseClassName(info.ClassInfoIndex),
				cp[desc.NameIndex].Value().(string), cp[desc.DescriptorIndex].Value().(string))
		}
	}
}

func (s *SimpleConsolePrinter) _printHeader() {
	fmt.Printf("Generic Info:\n")
	fmt.Printf("    %-25s %d\n", "Minor version:", s.classFile.MinorVersion())
	fmt.Printf("    %-25s %d\n", "Major version:", s.classFile.MajorVersion())
	fmt.Printf("    %-25s %d\n", "Constant pool size:", len(s.classFile.ConstantPool()))
	fmt.Printf("    %-25s  %d\n", "Access flags:", s.classFile.AccessFlags())
	fmt.Printf("    %-25s %s\n", "This class:", s._parseClassName(s.classFile.ThisClass()))
	fmt.Printf("    %-25s %s\n", "Super class:", s._parseClassName(s.classFile.SuperClass()))
	fmt.Printf("    %-25s %d\n", "Interfaces count:", len(s.classFile.InterfaceIndexSet()))
	fmt.Printf("    %-25s %d\n", "Fields count:", len(s.classFile.Fields()))
	fmt.Printf("    %-25s %d\n", "Methods count:", len(s.classFile.Methods()))
	fmt.Printf("    %-25s %d\n", "Attributes count:", len(s.classFile.Attributes()))
}

func (s *SimpleConsolePrinter) Print() {
	s._printHeader()
	s._printConstantPool()
}

func NewSimpleConsolePrinter(file *classfile.JvmClassFile) *SimpleConsolePrinter {
	return &SimpleConsolePrinter{
		classFile: file,
	}
}
