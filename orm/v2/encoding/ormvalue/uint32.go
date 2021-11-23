package ormvalue

import (
	"bytes"
	"encoding/binary"
	"io"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type Uint32Codec struct{}

func (u Uint32Codec) FixedSize() int {
	return 4
}

func (u Uint32Codec) Size(protoreflect.Value) (int, error) {
	return u.FixedSize(), nil
}

func (u Uint32Codec) IsOrdered() bool {
	return true
}

func (u Uint32Codec) Compare(v1, v2 protoreflect.Value) int {
	return compareUint(v1, v2)
}

func (u Uint32Codec) Decode(r *bytes.Reader) (protoreflect.Value, error) {
	var x uint32
	err := binary.Read(r, binary.BigEndian, &x)
	return protoreflect.ValueOfUint32(x), err
}

func (u Uint32Codec) Encode(value protoreflect.Value, w io.Writer) error {
	return binary.Write(w, binary.BigEndian, uint32(value.Uint()))
}
