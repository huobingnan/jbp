package classfile

import (
	"bytecodeparser/jvm/reader"
	"fmt"
	"math"
)

// DoubleConstant 双精度浮点数值常量池常量
type DoubleConstant struct {
	DoubleValue float64
}

func (d *DoubleConstant) Tag() int {
	return ConstantDoubleInfo
}

func (d *DoubleConstant) Value() interface{} {
	return d.DoubleValue
}

func (d *DoubleConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Double_info: %f>", d.DoubleValue)
}

func _newDoubleConstant(r *reader.ByteCodeReader) *DoubleConstant {
	if double, ok := r.ReadU8(); ok {
		return &DoubleConstant{
			DoubleValue: math.Float64frombits(double),
		}
	} else {
		panic("Read double constant error")
	}
}
