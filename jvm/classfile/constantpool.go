package classfile

import (
	"fmt"
	"github.com/huobingnan/jbp/jvm/reader"
)

const (
	ConstantUtf8Info               = 1  // UTF-8编码的字符串
	ConstantIntegerInfo            = 3  // 整型字面量
	ConstantFloatInfo              = 4  // 浮点型字面量
	ConstantLongInfo               = 5  // 长整型字面量
	ConstantDoubleInfo             = 6  // 双精度浮点型字面量
	ConstantClassInfo              = 7  // 类或者接口的符号引用
	ConstantStringInfo             = 8  // 字符串类型字面量
	ConstantFieldRefInfo           = 9  // 字段的符号引用
	ConstantMethodRefInfo          = 10 // 类中方法的符号引用
	ConstantInterfaceMethodRefInfo = 11 // 接口中方法的符号引用
	ConstantNameAndTypeInfo        = 12 // 字段或方法的部分符号引用
	ConstantMethodHandleInfo       = 15 // 表示方法句柄
	ConstantMethodTypeInfo         = 16 // 表示方法类型
	ConstantDynamicInfo            = 17 // 表示一个动态计算的常量
	ConstantInvokeDynamicInfo      = 18 // 表示一个动态方法调用点
	ConstantModuleInfo             = 19 // 表示一个模块
	ConstantPackageInfo            = 20 // 表示一个模块中开放或者导出的包
)

type ConstantSizeTable []int

func NewConstantSizeTable() ConstantSizeTable {
	table := make([]int, 21)
	table[ConstantUtf8Info] = 3
	table[ConstantIntegerInfo] = 4
	table[ConstantFloatInfo] = 4
	table[ConstantLongInfo] = 5
	table[ConstantDoubleInfo] = 8
	table[ConstantClassInfo] = 2
	table[ConstantStringInfo] = 2
	table[ConstantFieldRefInfo] = 4
	table[ConstantMethodRefInfo] = 4
	table[ConstantInterfaceMethodRefInfo] = 4
	table[ConstantNameAndTypeInfo] = 4
	table[ConstantMethodHandleInfo] = 3
	table[ConstantMethodTypeInfo] = 2
	table[ConstantDynamicInfo] = 4
	table[ConstantInvokeDynamicInfo] = 4
	table[ConstantModuleInfo] = 2
	table[ConstantPackageInfo] = 2
	return table
}

type JvmClassFileConstant interface {
	Tag() int
	Value() interface{}
}

type ConstantPool []JvmClassFileConstant

func NewConstant(r *reader.ByteCodeReader) JvmClassFileConstant {
	tag, ok := r.ReadU1()
	if !ok {
		panic("Read constant pool tag error")
	}
	switch tag {
	case ConstantUtf8Info:
		return _newUtf8Constant(r)
	case ConstantIntegerInfo:
		return _newIntegerConstant(r)
	case ConstantFloatInfo:
		return _newFloatConstant(r)
	case ConstantLongInfo:
		return _newLongConstant(r)
	case ConstantDoubleInfo:
		return _newDoubleConstant(r)
	case ConstantClassInfo:
		return _newClassConstant(r)
	case ConstantStringInfo:
		return _newStringConstant(r)
	case ConstantFieldRefInfo:
		return _newFieldRefConstant(r)
	case ConstantMethodRefInfo:
		return _newMethodRefConstant(r)
	case ConstantInterfaceMethodRefInfo:
		return _newInterfaceMethodRefConstant(r)
	case ConstantNameAndTypeInfo:
		return _newNameAndTypeConstant(r)
	case ConstantMethodHandleInfo:
		return _newMethodHandleConstant(r)
	case ConstantMethodTypeInfo:
		return _newMethodTypeConstant(r)
	case ConstantDynamicInfo:
		return _newDynamicConstant(r)
	case ConstantInvokeDynamicInfo:
		return _newInvokeDynamicConstant(r)
	case ConstantModuleInfo:
		return _newModuleConstant(r)
	case ConstantPackageInfo:
		return _newPackageConstant(r)
	default:
		panic(fmt.Sprintf("Error tag : %d", tag))
	}
}

// NewConstantPool 从字节码读取器中读取常量池的信息
// Notice: 当前字节码读取器的游标必须指向常量池开始的位置
func NewConstantPool(r *reader.ByteCodeReader) ConstantPool {
	size, ok := r.ReadU2()
	if !ok {
		panic("Can't read constant pool size")
	}
	constantPool := make([]JvmClassFileConstant, size)
	index := 1
	for i := 0; i < int(size-1); i++ {
		constantPool[index] = NewConstant(r)
		index += 1
	}
	return constantPool
}
