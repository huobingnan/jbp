package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

type ElementValuePair [2]uint16

func (e *ElementValuePair) Param() uint16 { return (*e)[0] }

func (e *ElementValuePair) Value() uint16 { return (*e)[1] }

type Annotation struct {
	TypeIndex         uint16
	ElementValuePairs []ElementValuePair
}

// RuntimeVisibleAnnotationsAttribute 是一个变长属性，它记录了类 、字段或方法的声明上记录运行时可见注解
type RuntimeVisibleAnnotationsAttribute struct {
	length      uint32
	Annotations []Annotation
}

func (r *RuntimeVisibleAnnotationsAttribute) Name() string { return RuntimeVisibleAnnotations }

func (r *RuntimeVisibleAnnotationsAttribute) Length() uint32 { return r.length }

// 读取注解信息
func readAnnotation(r *reader.ByteCodeReader) Annotation {
	var ok bool
	annotation := Annotation{}
	annotation.TypeIndex, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read runtime visible annotations attribute error",
			"can't read type_index info", r.Offset()))
	}
	elementPairsCount, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read runtime visible annotations attribute error",
			"can't read element_value_pairs_count info", r.Offset()))
	}
	annotation.ElementValuePairs = make([]ElementValuePair, 0, elementPairsCount)
	var t1, t2 uint16
	for i := uint16(0); i < elementPairsCount; i++ {
		t1, ok = r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read runtime visible annotations attribute error",
				"can't read element_value_pair_param info", r.Offset()))
		}
		t2, ok = r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read runtime visible annotations attribute error",
				"can't read element_value_pair_value info", r.Offset()))
		}
		annotation.ElementValuePairs = append(annotation.ElementValuePairs, [2]uint16{t1, t2})
	}
	return annotation
}

func _newRuntimeVisibleAnnotationsAttribute(r *reader.ByteCodeReader,
	cp constantpool.ConstantPool) *RuntimeVisibleAnnotationsAttribute {
	var ok bool
	ret := new(RuntimeVisibleAnnotationsAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read runtime visible annotations attribute error",
			"can't read Length info", r.Offset()))
	}
	numberOfAnnotation, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read runtime visible annotations attribute error",
			"can't read number_of_annotation info", r.Offset()))
	}
	ret.Annotations = make([]Annotation, 0, numberOfAnnotation)
	for i := uint16(0); i < numberOfAnnotation; i++ {
		ret.Annotations = append(ret.Annotations, readAnnotation(r))
	}
	return ret
}
