package main

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadConstantPool(s *reader.ByteCodeReader, size int) {
	constantSizeTable := constantpool.NewConstantSizeTable()
	for i := 0; i < size-1; i++ {
		tag, _ := s.ReadU1()
		if tag == constantpool.ConstantUtf8Info {
			utf8 := constantpool.NewUtf8Constant(s)
			fmt.Printf("%#v\n", utf8)
			continue
		}
		size := constantSizeTable[tag]
		switch tag {
		case constantpool.ConstantUtf8Info:
			fmt.Println("Utf8")
		case constantpool.ConstantIntegerInfo:
			i := constantpool.NewIntegerConstant(s)
			fmt.Printf("%#v\n", i)
		case constantpool.ConstantFloatInfo:
			fmt.Println("Float")
		case constantpool.ConstantLongInfo:
			fmt.Println("Long")
		case constantpool.ConstantDoubleInfo:
			fmt.Println("Double")
		case constantpool.ConstantClassInfo:
			fmt.Println("Class")
		case constantpool.ConstantStringInfo:
			fmt.Println("String")
		case constantpool.ConstantFieldRefInfo:
			fmt.Println("Field ref")
		case constantpool.ConstantMethodRefInfo:
			fmt.Println("Method ref")
		case constantpool.ConstantInterfaceMethodRefInfo:
			fmt.Println("Interface Method Ref")
		case constantpool.ConstantNameAndTypeInfo:
			nameAndType := constantpool.NewNameAndTypeConstant(s)
			fmt.Printf("%#v\n", nameAndType)
		case constantpool.ConstantMethodHandleInfo:
			fmt.Println("Method Handle")
		case constantpool.ConstantMethodTypeInfo:
			fmt.Println("Method Type")
		case constantpool.ConstantDynamicInfo:
			fmt.Println("Dynamic")
		case constantpool.ConstantInvokeDynamicInfo:
			fmt.Println("Invoke Dynamic")
		case constantpool.ConstantModuleInfo:
			fmt.Println("Module")
		case constantpool.ConstantPackageInfo:
			fmt.Println("Package")

		}
		s.ReadAny(size)
	}
}

func main() {
	fd, err := os.Open("/Users/huobingnan/Code/Golang/ByteCodeParser/data/BiLock.class")
	if err != nil {
		panic("open file error")
	}
	defer fd.Close()
	byteStream, err := ioutil.ReadAll(fd)
	if err != nil {
		panic("read file error")
	}
	stream := reader.NewByteCodeReader(byteStream)
	magic, _ := stream.ReadU4()
	fmt.Printf("%x", magic)
	minVersion, _ := stream.ReadU2()
	fmt.Printf(" %d", minVersion)
	majorVersion, _ := stream.ReadU2()
	fmt.Printf(" %d", majorVersion)
	constantPoolCnt, _ := stream.ReadU2()
	fmt.Printf(" %d\n", constantPoolCnt)
	ReadConstantPool(stream, int(constantPoolCnt))
}
