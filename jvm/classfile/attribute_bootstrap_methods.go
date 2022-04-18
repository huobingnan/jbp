package classfile

import (
	"github.com/huobingnan/jbp/jvm/reader"
)

type BootstrapMethod struct {
	MethodRef uint16   // 指向常量池中一个<CONSTANT_MethodHandle_info>
	Args      []uint16 // 有效地常量池索引
}

type BootstrapMethodsAttribute struct {
	length  uint32
	Methods []BootstrapMethod
}

func (b *BootstrapMethodsAttribute) Name() string { return BootstrapMethods }

func (b *BootstrapMethodsAttribute) Length() uint32 { return b.length }

func _newBootstrapMethodsAttribute(r *reader.ByteCodeReader, cp ConstantPool) *BootstrapMethodsAttribute {
	var ok bool
	ret := new(BootstrapMethodsAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read bootstrap methods attribute error",
			"can't read length info", r.Offset()))
	}
	numberOfBootstrapMethods, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read bootstrap methods attribute error",
			"can't read number_of_bootstrap_methods info", r.Offset()))
	}
	ret.Methods = make([]BootstrapMethod, 0, numberOfBootstrapMethods)
	for i := uint16(0); i < numberOfBootstrapMethods; i++ {
		// parse bootstrap_method info
		method := BootstrapMethod{}
		method.MethodRef, ok = r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read bootstrap methods attribute error",
				"can't read method_ref info", r.Offset()))
		}
		numberOfArgs, ok := r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read bootstrap methods attribute error",
				"can't read number_of_args info", r.Offset()))
		}
		method.Args = make([]uint16, 0, numberOfArgs)
		// parse args info
		for i := uint16(0); i < numberOfArgs; i++ {
			argIdx, ok := r.ReadU2()
			if !ok {
				panic(ErrorMsgFmt("Read bootstrap methods attribute error",
					"can't read arg_index info", r.Offset()))
			}
			method.Args = append(method.Args, argIdx)
		}
		ret.Methods = append(ret.Methods, method)
	}
	return ret
}
