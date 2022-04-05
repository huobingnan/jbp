package reader

import "encoding/binary"

type ByteCodeReader struct {
	buffer []byte
}

func (s *ByteCodeReader) ReadU1() (uint8, bool) {
	if len(s.buffer) <= 0 {
		return 0, false
	}
	ret := s.buffer[0]
	s.buffer = s.buffer[1:]
	return ret, true
}

func (s *ByteCodeReader) ReadU2() (uint16, bool) {
	if len(s.buffer) < 2 {
		return 0, false
	}
	ret := binary.BigEndian.Uint16(s.buffer[:2])
	s.buffer = s.buffer[2:]
	return ret, true
}

func (s *ByteCodeReader) ReadU4() (uint32, bool) {
	if len(s.buffer) < 4 {
		return 0, false
	}
	ret := binary.BigEndian.Uint32(s.buffer[:4])
	s.buffer = s.buffer[4:]
	return ret, true
}

func (s *ByteCodeReader) ReadU8() (uint64, bool) {
	if len(s.buffer) < 8 {
		return 0, false
	}
	ret := binary.BigEndian.Uint64(s.buffer[:8])
	s.buffer = s.buffer[8:]
	return ret, true
}

func (s *ByteCodeReader) ReadAny(size int) ([]byte, bool) {
	if size < 0 || len(s.buffer) < size {
		return nil, false
	}
	ret := s.buffer[:size]
	s.buffer = s.buffer[size:]
	return ret, true
}

func NewByteCodeReader(byteCode []byte) *ByteCodeReader {
	reader := new(ByteCodeReader)
	reader.buffer = byteCode
	return reader
}
