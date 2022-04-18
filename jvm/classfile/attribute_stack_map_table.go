package classfile

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
	"encoding/binary"
	"strconv"
)

const (
	SameFrame = iota
	SameLocals1StackItemFrame
	SameLocals1StackItemFrameExtended
	ChopFrame
	SameFrameExtended
	AppendFrame
	FullFrame
)

const (
	ItemTop = iota
	ItemInteger
	ItemFloat
	ItemDouble
	ItemLong
	ItemNull
	ItemUninitializedThis
	ItemObject
	ItemUninitialized
)

// StackMapFrame 这里选择不解析
type StackMapFrame struct {
	Source         []byte
	Type           int
	FrameTypeValue byte
}

type StackMapTableAttribute struct {
	length  uint32
	Entries []StackMapFrame
}

func (s *StackMapTableAttribute) Name() string { return StackMapTable }

func (s *StackMapTableAttribute) Length() uint32 { return 0 }

func _processVerificationTypeInfo(r *reader.ByteCodeReader, tag byte) []byte {
	switch {
	case tag == ItemObject || tag == ItemUninitialized:
		ret, ok := r.ReadAny(2)
		if !ok {
			panic(ErrorMsgFmt("Read stack map frame error", "can't read verification_type_info", r.Offset()))
		}
		return ret
	default:
		return nil
	}
}

func _readSameFrame(r *reader.ByteCodeReader, s *StackMapTableAttribute, value byte) {
	s.Entries = append(s.Entries, StackMapFrame{nil, SameFrame, value})
}

func _readSameLocals1StackItemFrame(r *reader.ByteCodeReader, s *StackMapTableAttribute, value byte) {
	tag, ok := r.ReadU1()
	if !ok {
		panic(ErrorMsgFmt("Read stack map frame `same_locals_1_stack_item_frame` error",
			strconv.Itoa(int(tag)), r.Offset()))
	}
	source := []byte{tag}
	verificationTypeInfo := _processVerificationTypeInfo(r, tag)
	if verificationTypeInfo != nil {
		for _, each := range verificationTypeInfo {
			source = append(source, each)
		}
	}
	s.Entries = append(s.Entries, StackMapFrame{source, SameLocals1StackItemFrame, value})
}

func _readSameLocals1StackItemFrameExtended(r *reader.ByteCodeReader, s *StackMapTableAttribute, value byte) {
	offset, ok := r.ReadAny(2)
	if !ok {
		panic(ErrorMsgFmt("Read stack map frame `same_locals_1_stack_item_frame_extended` error",
			"can't read offset", r.Offset()))
	}
	vtag, ok := r.ReadU1()
	if !ok {
		panic(ErrorMsgFmt("Read stack map frame `same_locals_1_stack_item_frame_extended` error",
			"can't read verification_tag info", r.Offset()))
	}
	source := []byte{vtag, offset[0], offset[1]}
	verificationTypeInfo := _processVerificationTypeInfo(r, vtag)
	if verificationTypeInfo != nil {
		for _, each := range verificationTypeInfo {
			source = append(source, each)
		}
	}
	s.Entries = append(s.Entries, StackMapFrame{source, SameLocals1StackItemFrameExtended, value})
}

func _readSameFrameExtended(r *reader.ByteCodeReader, s *StackMapTableAttribute, value byte) {
	offset, ok := r.ReadAny(2)
	if !ok {
		panic(ErrorMsgFmt("Read stack map frame `same_frame_extended` error",
			"can't read offset info", r.Offset()))
	}
	s.Entries = append(s.Entries, StackMapFrame{offset, SameFrameExtended, value})
}

func _readAppendFrame(r *reader.ByteCodeReader, s *StackMapTableAttribute, value byte) {
	offset, ok := r.ReadAny(2)
	if !ok {
		panic(ErrorMsgFmt("Read stack map frame `append_frame` error",
			"can't read offset info", r.Offset()))
	}
	source := []byte{offset[0], offset[1]}
	for i := uint8(0); i < value-251; i++ {
		vtag, ok := r.ReadU1()
		if !ok {
			panic(ErrorMsgFmt("Read stack map frame `append_frame` error",
				"can't read vtag info", r.Offset()))
		}
		verificationTypeInfo := _processVerificationTypeInfo(r, vtag)
		if verificationTypeInfo != nil {
			for _, each := range verificationTypeInfo {
				source = append(source, each)
			}
		}

	}
	s.Entries = append(s.Entries, StackMapFrame{source, ChopFrame, value})
}

func _readChopFrame(r *reader.ByteCodeReader, s *StackMapTableAttribute, value byte) {
	offset, ok := r.ReadAny(2)
	if !ok {
		panic(ErrorMsgFmt("Read stack map frame `chop_frame` error",
			"can't read offset info", r.Offset()))
	}
	s.Entries = append(s.Entries, StackMapFrame{offset, ChopFrame, value})
}

func _readFullFrame(r *reader.ByteCodeReader, s *StackMapTableAttribute, value byte) {
	var numbers []byte = make([]byte, 2) // 用于缓冲数字
	offset, ok := r.ReadAny(2)
	if !ok {
		panic(ErrorMsgFmt("Read stack map frame `full_frame` error",
			"can't read offset info", r.Offset()))
	}
	numberOfLocals, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read stack map frame `chop_frame` error",
			"can't read number_of_locals info", r.Offset()))
	}
	source := []byte{offset[0], offset[1]}
	binary.BigEndian.PutUint16(numbers, numberOfLocals)
	for _, each := range numbers {
		source = append(source, each)
	}
	for i := uint16(0); i < numberOfLocals; i++ {
		vtag, ok := r.ReadU1()
		if !ok {
			panic(ErrorMsgFmt("Read stack map frame `chop_frame` error",
				"can't read locals_tag info", r.Offset()))
		}
		verificationTypeInfo := _processVerificationTypeInfo(r, vtag)
		if verificationTypeInfo != nil {
			for _, each := range verificationTypeInfo {
				source = append(source, each)
			}
		}
	}
	numberOfStackItem, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read stack map frame `chop_frame` error",
			"can't read number_of_stack_item info", r.Offset()))
	}
	binary.BigEndian.PutUint16(numbers, numberOfStackItem)
	for _, each := range numbers {
		source = append(source, each)
	}
	for i := uint16(0); i < numberOfStackItem; i++ {
		vtag, ok := r.ReadU1()
		if !ok {
			panic(ErrorMsgFmt("Read stack map frame `chop_frame` error",
				"can't read stack_item info", r.Offset()))
		}
		verificationTypeInfo := _processVerificationTypeInfo(r, vtag)
		if verificationTypeInfo != nil {
			for _, each := range verificationTypeInfo {
				source = append(source, each)
			}
		}
	}
	s.Entries = append(s.Entries, StackMapFrame{source, FullFrame, value})
}

func _newStackMapTableAttribute(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *StackMapTableAttribute {
	var ok bool
	ret := new(StackMapTableAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read stack map Table attribute error",
			"can't read Length info", r.Offset()))
	}
	numberOfEntries, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read stack map Table attribute error",
			"can't read number_of_entries info", r.Offset()))
	}

	ret.Entries = make([]StackMapFrame, 0, numberOfEntries)
	for i := uint16(0); i < numberOfEntries; i++ {
		frameType, ok := r.ReadU1()
		if !ok {
			panic(ErrorMsgFmt("Read stack map Table attribute error",
				"can't read frame_type info", r.Offset()))
		}
		if frameType <= 63 {
			_readSameFrame(r, ret, frameType)
		} else if 64 <= frameType && frameType <= 127 {
			_readSameLocals1StackItemFrame(r, ret, frameType)
		} else if frameType == 247 {
			_readSameLocals1StackItemFrameExtended(r, ret, frameType)
		} else if 248 <= frameType && frameType <= 250 {
			_readChopFrame(r, ret, frameType)
		} else if frameType == 251 {
			_readSameFrameExtended(r, ret, frameType)
		} else if 252 <= frameType && frameType <= 254 {
			_readAppendFrame(r, ret, frameType)
		} else {
			_readFullFrame(r, ret, frameType)
		}
	}
	return ret
}
