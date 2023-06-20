package evm

import (
	"testing"

	util "github.com/TafveezA/EVM/util"
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
	data := []byte{0x03, 0x0a, 0x46, 0x0c, 0x4f, 0x0c, 0x4f, 0x0c, 0x0d, 0x05, 0x0a, 0x0f}
	contractState := NewState()
	vm := NewVM(data, contractState)
	assert.Nil(t, vm.Run())

	valueBytes, err := contractState.Get([]byte("FOO"))
	value := util.DeserialzeInt64(valueBytes)
	assert.Nil(t, err)
	assert.Equal(t, value, int64(5))

}
func TestVMMul(t *testing.T) {
	data := []byte{0x02, 0x0a, 0x02, 0x0a, 0xea}
	contractState := NewState()
	vm := NewVM(data, contractState)
	assert.Nil(t, vm.Run())
	result := vm.stack.Pop().(int)
	assert.Equal(t, result, 4)

}

func TestVMDiv(t *testing.T) {
	data := []byte{0x04, 0x0a, 0x02, 0x0a, 0xfd}
	contractState := NewState()
	vm := NewVM(data, contractState)
	assert.Nil(t, vm.Run())
	result := vm.stack.Pop().(int)
	assert.Equal(t, result, 2)

}
