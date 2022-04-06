package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

const (
	LocalVariableTableContent = "content"
)

type LocalVariableInfo struct {
	// start_pc和length属性分别代表了这个局部变量的生命周期开始的字节码偏移量及其作用范围覆盖的长度
	// 两者结合起来就是这个局部变量在字节码之中的作用域范围
	startPc uint16
	length  uint16
	// name_index和descriptor_index都是指向常量池中CONSTANT_Utf8_info型常量的索引
	// 分别代表了局部变量的名称以及这个局部变量的描述符
	nameIndex       uint16
	descriptorIndex uint16
	// index是这个局部变量在栈帧的局部变量表中变量槽的位置
	// 当这个变量数据类型是64位类型时 (double和long)，它占用的变量槽为index和index+1两个。
	index uint16
}

type LocalVariableTableAttribute struct {
	length uint32
	table  []LocalVariableInfo
}

func (self *LocalVariableTableAttribute) Name() string { return LocalVariableTable }

func (self *LocalVariableTableAttribute) Length() uint32 { return self.length }

func (self *LocalVariableTableAttribute) Get(key string) interface{} {
	switch key {
	case LocalVariableTableContent:
		return self.table
	default:
		return nil
	}
}

func readLocalVariableInfo(r *reader.ByteCodeReader) LocalVariableInfo {
	info := LocalVariableInfo{}
	var ok bool
	info.startPc, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error", "can't read start_pc info", r.Offset()))
	}
	info.length, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error", "can't read length info", r.Offset()))
	}
	info.nameIndex, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error", "can't read name_index info", r.Offset()))
	}
	info.descriptorIndex, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error",
			"can't read descriptor_index info", r.Offset()))
	}
	info.index, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error", "can't read index info", r.Offset()))
	}
	return info
}

func NewLocalVariableTableAttribute(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *LocalVariableTableAttribute {
	var ok bool
	ret := new(LocalVariableTableAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error", "can't read length info", r.Offset()))
	}
	tableLength, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error", "can't read table_length info", r.Offset()))
	}
	ret.table = make([]LocalVariableInfo, 0, tableLength)
	for i := uint16(0); i < tableLength; i++ {
		ret.table = append(ret.table, readLocalVariableInfo(r))
	}
	return ret
}
