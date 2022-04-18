package classfile

import (
	"bytecodeparser/jvm/reader"
)

type NestMembersAttribute struct {
	length       uint32
	ClassesIndex []uint16 // 类在常量池中的索引
}

func (n *NestMembersAttribute) Name() string { return NestMembers }

func (n *NestMembersAttribute) Length() uint32 { return n.length }

func _newNestMembersAttribute(r *reader.ByteCodeReader, cp ConstantPool) *NestMembersAttribute {
	var ok bool
	ret := new(NestMembersAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read nest members attribute error",
			"can't read length info", r.Offset()))
	}
	numberOfClasses, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read nest members attribute error",
			"can't read number_of_classes info", r.Offset()))
	}
	ret.ClassesIndex = make([]uint16, 0, numberOfClasses)
	for i := uint16(0); i < numberOfClasses; i++ {
		idx, ok := r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read nest members attribute error",
				"can't read class_index info", r.Offset()))
		}
		ret.ClassesIndex = append(ret.ClassesIndex, idx)
	}
	return ret
}
