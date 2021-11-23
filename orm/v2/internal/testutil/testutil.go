package testutil

import (
	"fmt"
	"strings"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormkey"

	"github.com/regen-network/regen-ledger/orm/v2/encoding/ormvalue"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"

	"github.com/regen-network/regen-ledger/orm/v2/internal/testpb"
	"google.golang.org/protobuf/reflect/protoreflect"
	"pgregory.net/rapid"
)

type TestKeyPartSpec struct {
	FieldName string
	Gen       *rapid.Generator
}

var TestKeyPartSpecs = []TestKeyPartSpec{
	{
		"UINT32",
		rapid.Uint32(),
	},
	{
		"UINT64",
		rapid.Uint64(),
	},
	{
		"STRING",
		rapid.String().Filter(func(x string) bool {
			// filter out null terminators
			return strings.IndexByte(x, 0) < 0
		}),
	},
	{
		"BYTES",
		rapid.SliceOfN(rapid.Byte(), 0, 255),
	},
}

func MakeTestPartCodec(fname string, nonTerminal bool) (ormvalue.Codec, error) {
	return ormvalue.MakeCodec(GetTestField(fname), nonTerminal)
}

func GetTestField(fname string) protoreflect.FieldDescriptor {
	a := &testpb.A{}
	return ormkey.GetFieldDescriptor(a.ProtoReflect().Descriptor(), fname)
}

type TestKey struct {
	KeySpecs []TestKeyPartSpec
	Fields   string
	Codec    *ormkey.Codec
}

var TestKeyGen = rapid.SliceOfN(rapid.IntRange(0, len(TestKeyPartSpecs)-1), 1, len(TestKeyPartSpecs)).
	Filter(func(xs []int) bool {
		have := map[int]bool{}
		for _, x := range xs {
			if have[x] {
				return false
			}
			have[x] = true
		}
		return true
	}).Map(func(xs []int) TestKey {
	var specs []TestKeyPartSpec
	var fields []protoreflect.FieldDescriptor
	var fnames []string

	for _, x := range xs {
		spec := TestKeyPartSpecs[x]
		specs = append(specs, spec)
		fields = append(fields, GetTestField(spec.FieldName))
		fnames = append(fnames, spec.FieldName)
	}

	cdc, err := ormkey.MakeCodec([]byte{1}, fields)
	if err != nil {
		panic(err)
	}

	return TestKey{
		Codec:    cdc,
		KeySpecs: specs,
		Fields:   strings.Join(fnames, ","),
	}
},
)

func (k TestKey) Draw(t *rapid.T, id string) []protoreflect.Value {
	n := len(k.KeySpecs)
	keyValues := make([]protoreflect.Value, n)
	for i, k := range k.KeySpecs {
		keyValues[i] = protoreflect.ValueOf(k.Gen.Draw(t, fmt.Sprintf("%s[%d]", id, i)))
	}
	return keyValues
}

func (k TestKey) RequireValuesEqual(t require.TestingT, values, values2 []protoreflect.Value) {
	for i := 0; i < len(values); i++ {
		assert.Equal(t, 0, k.Codec.ValueCodecs[i].Compare(values[i], values2[i]),
			"values[%d]: %v != %v", i, values[i].Interface(), values2[i].Interface())
	}
}
