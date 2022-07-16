package bind

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type (
	queryTestStruct struct {
		I           int
		PtrI        *int
		I8          int8
		PtrI8       *int8
		I16         int16
		PtrI16      *int16
		I32         int32
		PtrI32      *int32
		I64         int64
		PtrI64      *int64
		UI          uint
		PtrUI       *uint
		UI8         uint8
		PtrUI8      *uint8
		UI16        uint16
		PtrUI16     *uint16
		UI32        uint32
		PtrUI32     *uint32
		UI64        uint64
		PtrUI64     *uint64
		B           bool
		PtrB        *bool
		F32         float32
		PtrF32      *float32
		F64         float64
		PtrF64      *float64
		S           string
		PtrS        *string
		cantSet     string
		DoesntExist string
		GoT         time.Time
		GoTptr      *time.Time
		T           Timestamp
		Tptr        *Timestamp
		SA          StringArray
	}

	queryTestStructWithTags struct {
		I           int      `query:"I"`
		PtrI        *int     `query:"PtrI"`
		I8          int8     `query:"I8"`
		PtrI8       *int8    `query:"PtrI8"`
		I16         int16    `query:"I16"`
		PtrI16      *int16   `query:"PtrI16"`
		I32         int32    `query:"I32"`
		PtrI32      *int32   `query:"PtrI32"`
		I64         int64    `query:"I64"`
		PtrI64      *int64   `query:"PtrI64"`
		UI          uint     `query:"UI"`
		PtrUI       *uint    `query:"PtrUI"`
		UI8         uint8    `query:"UI8"`
		PtrUI8      *uint8   `query:"PtrUI8"`
		UI16        uint16   `query:"UI16"`
		PtrUI16     *uint16  `query:"PtrUI16"`
		UI32        uint32   `query:"UI32"`
		PtrUI32     *uint32  `query:"PtrUI32"`
		UI64        uint64   `query:"UI64"`
		PtrUI64     *uint64  `query:"PtrUI64"`
		B           bool     `query:"B"`
		PtrB        *bool    `query:"PtrB"`
		F32         float32  `query:"F32"`
		PtrF32      *float32 `query:"PtrF32"`
		F64         float64  `query:"F64"`
		PtrF64      *float64 `query:"PtrF64"`
		S           string   `query:"S"`
		PtrS        *string  `query:"PtrS"`
		cantSet     string
		DoesntExist string      `query:"DoesntExist"`
		GoT         time.Time   `query:"GoT"`
		GoTptr      *time.Time  `query:"GoTptr"`
		T           Timestamp   `query:"T"`
		Tptr        *Timestamp  `query:"Tptr"`
		SA          StringArray `query:"SA"`
	}

	Timestamp   time.Time
	TimeArray   []Timestamp
	StringArray []string
	Struct      struct {
		Foo string
	}
	Bar struct {
		Baz int `query:"baz"`
	}
)

func TestQuery_NilData(t *testing.T) {
	receiver := struct {
		Key string `query:"key"`
	}{}

	err := Query(&receiver, nil)
	assert.Nil(t, err)
	assert.Empty(t, receiver)
}

func TestQuery_NilReceiver(t *testing.T) {
	data := map[string][]string{}
	data["key"] = []string{"value"}

	err := Query(nil, data)
	assert.Nil(t, err)
}

func TestQuery_MapReceiver(t *testing.T) {
	receiver := map[string]string{}
	receiver["key"] = ""

	data := map[string][]string{}
	data["key"] = []string{"value"}

	err := Query(&receiver, data)
	assert.Nil(t, err)
	assert.Equal(t, "value", receiver["key"])
}

func TestQuery_InvalidReceiver(t *testing.T) {
	receiver := 1
	data := map[string][]string{}
	data["key"] = []string{"value"}

	err := Query(&receiver, data)
	assert.NotNil(t, err)
}

func TestEnv(t *testing.T) {
	os.Setenv("BIND_TEST_VALUE", "23")

	type Example struct {
		Value int `env:"BIND_TEST_VALUE"`
	}

	var e Example

	err := Env(&e)
	require.Nil(t, err)
	assert.Equal(t, 23, e.Value)

	os.Unsetenv("BIND_TEST_VALUE")
}

func TestHeader(t *testing.T) {
	receiver := struct {
		Key string `header:"X-Key"`
	}{}

	data := map[string][]string{}
	data["X-Key"] = []string{"value"}

	err := Header(&receiver, data)
	require.Nil(t, err)
	assert.Equal(t, "value", receiver.Key)
}
