package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
	"math"
)

// DoubleConstant 双精度浮点数值常量池常量
type DoubleConstant struct {
	value float64
}

func (d *DoubleConstant) Tag() int {
	return ConstantDoubleInfo
}

func (d *DoubleConstant) Value() interface{} {
	return d.value
}

func (d *DoubleConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Double_info: %f>", d.value)
}

func NewDoubleConstant(r *reader.ByteCodeReader) *DoubleConstant {
	if double, ok := r.ReadU8(); ok {
		return &DoubleConstant{
			value: math.Float64frombits(double),
		}
	} else {
		panic("Read double constant error")
	}
}
