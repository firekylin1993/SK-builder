package kits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommonFunc(t *testing.T) {
	t.Run("BytesCombine", func(t *testing.T) {
		bit1 := []byte("hello")
		bit2 := []byte("world")
		combine := BytesCombine(bit1, bit2)
		assert.Equal(t, []byte("helloworld"), combine)
	})

	t.Run("IntToBytes", func(t *testing.T) {
		var myint uint32 = 123
		bytes := IntToBytes(myint)
		assert.IsType(t, []byte{}, bytes)
	})

	t.Run("GetCRC32Key", func(t *testing.T) {
		bit1 := []byte("hello")
		key := GetCRC32Key(bit1)
		assert.Equal(t, uint32(907060870), key)
	})
}
