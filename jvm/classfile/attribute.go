package classfile

import (
	"bytecodeparser/jvm/classfile/reader"
)

const (
	Code                      = "Code"                      // 代码属性
	ConstantValue             = "ConstantValue"             // 常量
	Deprecated                = "Deprecated"                // 废弃的字段
	Exceptions                = "Exceptions"                // 异常
	LineNumberTable           = "LineNumberTable"           // 行数表
	LocalVariableTable        = "LocalVariableTable"        // 本地变量表
	SourceFile                = "SourceFile"                // 源文件
	Synthetic                 = "Synthetic"                 // 非用户代码生成
	ExceptionTable            = "Exception Table"           // 异常表
	StackMapTable             = "StackMapTable"             // 栈映射表
	InnerClasses              = "InnerClasses"              // 内部类
	RuntimeVisibleAnnotations = "RuntimeVisibleAnnotations" // 运行时可见注解
	BootstrapMethods          = "BootstrapMethods"          // 启动方法
	NestMembers               = "NestMembers"               // 内部成员
)

// JvmClassFileAttribute 属性接口定义
type JvmClassFileAttribute interface {
	Name() string   // 获取属性名
	Length() uint32 // 获取属性的长度
}

func NewAttribute(r *reader.ByteCodeReader, cp ConstantPool) JvmClassFileAttribute {
	var name string
	if idx, ok := r.ReadU2(); ok {
		name = cp[idx].Value().(string)
	} else {
		panic(ErrorMsgFmt("Read attribute error", "can't read attribute_name_index info", r.Offset()))
	}
	switch name {
	case Code:
		return _newCodeAttribute(r, cp)
	case ConstantValue:
		return _newConstantValueAttribute(r, cp)
	case Deprecated:
		return _newDeprecatedAttribute(r, cp)
	case Exceptions:
		return _newExceptionsAttribute(r, cp)
	case LineNumberTable:
		return _newLineNumberTableAttribute(r, cp)
	case LocalVariableTable:
		return _newLocalVariableTableAttribute(r, cp)
	case SourceFile:
		return _newSourceFileAttribute(r, cp)
	case Synthetic:
		return _newSyntheticAttribute(r, cp)
	case StackMapTable:
		return _newStackMapTableAttribute(r, cp)
	case RuntimeVisibleAnnotations:
		return _newRuntimeVisibleAnnotationsAttribute(r, cp)
	case InnerClasses:
		return _newInnerClassAttribute(r, cp)
	case BootstrapMethods:
		return _newBootstrapMethodsAttribute(r, cp)
	case NestMembers:
		return _newNestMembersAttribute(r, cp)
	default:
		panic(ErrorMsgFmt("Unsupported attribute", name, r.Offset()))
	}
}
