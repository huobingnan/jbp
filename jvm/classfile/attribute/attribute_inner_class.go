package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

type InnerClassInfo struct {
	InnerClassIndex       uint16
	OuterClassIndex       uint16
	InnerNameIndex        uint16
	InnerClassAccessFlags uint16
}

type InnerClassAttribute struct {
	length  uint32
	Classes []InnerClassInfo
}

func (i *InnerClassAttribute) Name() string { return InnerClasses }

func (i *InnerClassAttribute) Length() uint32 { return i.length }

func NewInnerClassAttribute(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *InnerClassAttribute {
	var ok bool
	ret := new(InnerClassAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read inner class attribute error", "can't read length info", r.Offset()))
	}
	numberOfClasses, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read inner class attribute error",
			"can't read number_of_classes info", r.Offset()))
	}
	ret.Classes = make([]InnerClassInfo, 0, numberOfClasses)
	for i := uint16(0); i < numberOfClasses; i++ {
		// read inner class info
		class := InnerClassInfo{}
		class.InnerClassIndex, ok = r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read inner class attribute error",
				"can't read inner_class_info_index info", r.Offset()))
		}
		class.OuterClassIndex, ok = r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read inner class attribute error",
				"can't read outer_class_info_index info", r.Offset()))
		}
		class.InnerNameIndex, ok = r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read inner class attribute error",
				"can't read inner_class_name info", r.Offset()))
		}
		class.InnerClassAccessFlags, ok = r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read inner class attribute error",
				"can't read inner_class_access_flags info", r.Offset()))
		}
		ret.Classes = append(ret.Classes, class)
	}
	return ret
}
