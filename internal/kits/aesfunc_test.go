package kits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAes(t *testing.T) {
	b, err := AesEncrypt([]byte("i am a worker man"), []byte("key1234567890123"))
	assert.NoError(t, err)

	r, err := AesDecrypt(b, []byte("key1234567890123"))
	assert.NoError(t, err)
	fmt.Println(string(r))
}
