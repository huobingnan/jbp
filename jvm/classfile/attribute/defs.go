package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

const (
	Code               = "Code"               // 代码属性
	ConstantValue      = "ConstantValue"      // 常量
	Deprecated         = "Deprecated"         // 废弃的字段
	Exceptions         = "Exceptions"         // 异常
	LineNumberTable    = "LineNumberTable"    // 行数表
	LocalVariableTable = "LocalVariableTable" // 本地变量表
	SourceFile         = "SourceFile"         // 源文件
	Synthetic          = "Synthetic"          // 非用户代码生成
	ExceptionTable     = "Exception table"    // 异常表
	StackMapTable      = "StackMapTable"      // 栈映射表
	InnerClass         = "InnerClass"         // 内部类
)

// Attribute 属性接口定义
type Attribute interface {
	Name() string           // 获取属性名
	Length() uint32         // 获取属性的长度
	Get(string) interface{} // 获取属性
}

func New(r *reader.ByteCodeReader, cp constantpool.ConstantPool) Attribute {
	var name string
	if idx, ok := r.ReadU2(); ok {
		name = cp[idx].Value().(string)
	} else {
		panic("Read attribute error (name index)")
	}
	switch name {
	case Code:
		return NewCodeAttribute(r, cp)
	case ConstantValue:
		return NewConstantValueAttribute(r, cp)
	case Deprecated:
		return NewDeprecatedAttribute(r, cp)
	case Exceptions:
		return NewExceptionsAttribute(r, cp)
	case LineNumberTable:
		return NewLineNumberTableAttribute(r, cp)
	case LocalVariableTable:
		return NewLocalVariableTableAttribute(r, cp)
	case SourceFile:
		return NewSourceFileAttribute(r, cp)
	case Synthetic:
		return NewSyntheticAttribute(r, cp)
	default:
		panic(fmt.Sprintf("Unsupported attribute (%s)", name))
	}
}
