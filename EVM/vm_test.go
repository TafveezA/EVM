package evm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack(128)
	s.Push(1)
	s.Push(2)

	value := s.Pop()
	assert.Equal(t, value, 1)

	value = s.Pop()
	assert.Equal(t, value, 2)
}

func TestVM(t *testing.T) {
	pushFoo := []byte{0x4f, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d}
	data := []byte{0x02, 0x0a, 0x03, 0x0a, 0x4f, 0x0c, 0x4f, 0x0c, 0x46, 0x0c, 0x03, 0x0a, 0x0d, 0x0f}
	data = append(data, pushFoo...)
	contractState := NewState()

	vm := NewVM(data, contractState)
	assert.Nil(t, vm.Run())
	fmt.Printf("%v\n", vm.stack.data)
	// valueBytes, err := contractState.Get([]byte("FOO"))
	// value := util.DeserialzeInt64(valueBytes)
	// assert.Nil(t, err)
	// assert.Equal(t, value, int64(5))

	//result := vm.stack.Pop().([]byte)
	//assert.Equal(t, "FOO", string(result))

	// fmt.Printf("%v\n", contractState)

}
