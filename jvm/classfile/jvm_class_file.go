package classfile

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

const (
	ACC_PUBLIC     = 0x0001 // 标识类公开
	ACC_PRIVATE    = 0x0002 // 标识类私有
	ACC_PROTECTED  = 0x0004 // 标识类保护
	ACC_STATIC     = 0x0008 // 标识静态
	ACC_FINAL      = 0x0010 // 是否是final类
	ACC_SUPER      = 0x0020 // 是否允许使用invokespecial
	ACC_VOLATILE   = 0x0040 // 标识volatile
	ACC_TRANSIENT  = 0x0080 // 标识不可序列化
	ACC_INTERFACE  = 0x0200 // 标识一个接口类
	ACC_ABSTRACT   = 0x0400 // 标识抽象类
	ACC_SYNTHETIC  = 0x1000 // 标识这个类并非由用户代码产生
	ACC_ANNOTATION = 0x2000 // 标识这是一个注解
	ACC_ENUM       = 0x4000 // 标识这是一个枚举
	ACC_MODULE     = 0x8000 // 标识这是一个模块
)

type JvmClassFile struct {
	magicNumber       uint32                    // 魔数
	majorVersion      uint16                    // 主版本号
	minorVersion      uint16                    // 次版本号
	cp                constantpool.ConstantPool // 常量池
	accessFlags       uint16                    // 访问标志
	thisClass         uint16                    // 此类全限定名的常量池索引
	superClass        uint16                    // 此类父类的全限定名的常量池索引
	interfaceIndexSet []uint16                  // 此类实现的接口索引集合
}

// 读取接口索引集合
func readInterfaceSet(r *reader.ByteCodeReader, cfile *JvmClassFile) {
	size, ok := r.ReadU2()
	if !ok {
		panic("Read interface set size error")
	}
	if size == 0 {
		return
	}
	cfile.interfaceIndexSet = make([]uint16, size)
	for i := 0; i < int(size); i++ {
		if idx, ok := r.ReadU2(); ok {
			cfile.interfaceIndexSet[i] = idx
		} else {
			panic("Read interface index error")
		}
	}
}

func NewJvmClassFile(r *reader.ByteCodeReader) *JvmClassFile {
	ret := new(JvmClassFile)
	if magicNumber, ok := r.ReadU4(); ok && magicNumber == 0xcafebabe {
		ret.magicNumber = magicNumber
	} else {
		panic("Magic number validation error")
	}
	if minorVersion, ok := r.ReadU2(); ok {
		ret.minorVersion = minorVersion
	} else {
		panic("Read minor version error")
	}
	if majorVersion, ok := r.ReadU2(); ok {
		ret.majorVersion = majorVersion
	} else {
		panic("Read major version error")
	}
	// 读取常量池
	ret.cp = constantpool.NewConstantPool(r)
	// 读取访问标志
	if flags, ok := r.ReadU2(); ok {
		ret.accessFlags = flags
	} else {
		panic("Read access flags error")
	}
	// 读取本类的全限定类名
	if index, ok := r.ReadU2(); ok {
		ret.thisClass = index
	} else {
		panic("Read this class error")
	}
	// 读取父类的全限定类名
	if index, ok := r.ReadU2(); ok {
		ret.superClass = index
	} else {
		panic("Read super class error")
	}
	// 读取接口索引集合
	readInterfaceSet(r, ret)
	return ret
}

func (j *JvmClassFile) MagicNumber() uint32 { return j.magicNumber }

func (j *JvmClassFile) MajorVersion() uint16 { return j.majorVersion }

func (j *JvmClassFile) MinorVersion() uint16 { return j.minorVersion }

func (j *JvmClassFile) ConstantPool() constantpool.ConstantPool { return j.cp }

func (j *JvmClassFile) AccessFlags() uint16 { return j.accessFlags }

func (j *JvmClassFile) InterfaceIndexSet() []uint16 { return j.interfaceIndexSet }
