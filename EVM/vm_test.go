package evm

import (
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

	data := []byte{0x02, 0x0a, 0x61, 0x0c, 0x61, 0x0c, 0x0d}
	vm := NewVM(data)
	assert.Nil(t, vm.Run())
	result := vm.stack.Pop()
	assert.Equal(t, 4, result)

}
