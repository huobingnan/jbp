package classfile

import (
	"fmt"
	"github.com/huobingnan/jbp/jvm/reader"
)

const (
	AccPUBLIC     = 0x0001 // 标识类公开
	AccPRIVATE    = 0x0002 // 标识类私有
	AccPROTECTED  = 0x0004 // 标识类保护
	AccSTATIC     = 0x0008 // 标识静态
	AccFINAL      = 0x0010 // 是否是final类
	AccSUPER      = 0x0020 // 是否允许使用invoke special
	AccVOLATILE   = 0x0040 // 标识volatile
	AccTRANSIENT  = 0x0080 // 标识不可序列化
	AccINTERFACE  = 0x0200 // 标识一个接口类
	AccABSTRACT   = 0x0400 // 标识抽象类
	AccSYNTHETIC  = 0x1000 // 标识这个类并非由用户代码产生
	AccANNOTATION = 0x2000 // 标识这是一个注解
	AccENUM       = 0x4000 // 标识这是一个枚举
	AccMODULE     = 0x8000 // 标识这是一个模块
)

func ErrorMsgFmt(body, detail string, offset uint32) string {
	return fmt.Sprintf("%s (%s) @%d", body, detail, offset)
}

type JvmClassFile struct {
	magicNumber       uint32                  // 魔数
	majorVersion      uint16                  // 主版本号
	minorVersion      uint16                  // 次版本号
	cp                ConstantPool            // 常量池
	accessFlags       uint16                  // 访问标志
	thisClass         uint16                  // 此类全限定名的常量池索引
	superClass        uint16                  // 此类父类的全限定名的常量池索引
	interfaceIndexSet []uint16                // 此类实现的接口索引集合
	fields            []JvmClassFileField     // 字段表
	methods           []JvmClassFileMethod    // 方法表
	attributes        []JvmClassFileAttribute // 属性表
}

// 读取接口索引集合
func readInterfaceSet(r *reader.ByteCodeReader, cfile *JvmClassFile) {
	size, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read interface set error", "can't read size info", r.Offset()))
	}
	if size == 0 {
		return
	}
	cfile.interfaceIndexSet = make([]uint16, size)
	for i := 0; i < int(size); i++ {
		if idx, ok := r.ReadU2(); ok {
			cfile.interfaceIndexSet[i] = idx
		} else {
			panic(ErrorMsgFmt("Read interface set error",
				"can't read interface_index info", r.Offset()))
		}
	}
}

// 读取字段表
func readFieldTable(r *reader.ByteCodeReader, cfile *JvmClassFile) {
	length, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read field length error", "fatal", r.Offset()))
	}
	cfile.fields = make([]JvmClassFileField, 0, length)
	for i := uint16(0); i < length; i++ {
		cfile.fields = append(cfile.fields, *NewField(r, cfile.cp))
	}
}

// 读取魔数和版本信息
func readMagicNumberAndVersion(r *reader.ByteCodeReader, cfile *JvmClassFile) {
	if magicNumber, ok := r.ReadU4(); ok && magicNumber == 0xcafebabe {
		cfile.magicNumber = magicNumber
	} else {
		panic(ErrorMsgFmt("Read magic number error", "fatal", r.Offset()))
	}
	if minorVersion, ok := r.ReadU2(); ok {
		cfile.minorVersion = minorVersion
	} else {
		panic(ErrorMsgFmt("Read minor version error", "fatal", r.Offset()))
	}
	if majorVersion, ok := r.ReadU2(); ok {
		cfile.majorVersion = majorVersion
	} else {
		panic(ErrorMsgFmt("Read major version error", "fatal", r.Offset()))
	}
}

// 读取方法表
func readMethodTable(r *reader.ByteCodeReader, cfile *JvmClassFile) {
	length, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read method error", "fatal", r.Offset()))
	}
	cfile.methods = make([]JvmClassFileMethod, 0, length)
	for i := uint16(0); i < length; i++ {
		cfile.methods = append(cfile.methods, *NewMethod(r, cfile.cp))
	}
}

// 读取属性表
func _readAttributeTable(r *reader.ByteCodeReader, cfile *JvmClassFile) {
	attrLen, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read attribute table error", "fatal", r.Offset()))
	}
	cfile.attributes = make([]JvmClassFileAttribute, 0, attrLen)
	for i := uint16(0); i < attrLen; i++ {
		cfile.attributes = append(cfile.attributes, NewAttribute(r, cfile.cp))
	}
}

func NewJvmClassFile(r *reader.ByteCodeReader) *JvmClassFile {
	ret := new(JvmClassFile)
	readMagicNumberAndVersion(r, ret)
	ret.cp = NewConstantPool(r)
	// 读取访问标志
	if flags, ok := r.ReadU2(); ok {
		ret.accessFlags = flags
	} else {
		panic(ErrorMsgFmt("Read access flags error", "fatal", r.Offset()))
	}
	// 读取本类的全限定类名
	if index, ok := r.ReadU2(); ok {
		ret.thisClass = index
	} else {
		panic(ErrorMsgFmt("Read this class index error", "fatal", r.Offset()))
	}
	// 读取父类的全限定类名
	if index, ok := r.ReadU2(); ok {
		ret.superClass = index
	} else {
		panic(ErrorMsgFmt("Read super class index error", "fatal", r.Offset()))
	}
	readInterfaceSet(r, ret)
	readFieldTable(r, ret)
	readMethodTable(r, ret)
	_readAttributeTable(r, ret)
	return ret
}

func (j *JvmClassFile) MagicNumber() uint32 { return j.magicNumber }

func (j *JvmClassFile) MajorVersion() uint16 { return j.majorVersion }

func (j *JvmClassFile) MinorVersion() uint16 { return j.minorVersion }

func (j *JvmClassFile) ConstantPool() ConstantPool { return j.cp }

func (j *JvmClassFile) AccessFlags() uint16 { return j.accessFlags }

func (j *JvmClassFile) InterfaceIndexSet() []uint16 { return j.interfaceIndexSet }

func (j *JvmClassFile) Fields() []JvmClassFileField { return j.fields }

func (j *JvmClassFile) Methods() []JvmClassFileMethod { return j.methods }

func (j *JvmClassFile) Attributes() []JvmClassFileAttribute { return j.attributes }

func (j *JvmClassFile) ThisClass() uint16 { return j.thisClass }

func (j *JvmClassFile) SuperClass() uint16 { return j.superClass }
