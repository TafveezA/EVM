package evm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *assert.TestingT) {
	s := NewStack(128)
	s.Push(1)
	s.Push(2)

	value := s.Pop()
	assert.Equal(*t, value, 1)
	fmt.Println(s.data)
}

func TestVM(t *testing.T) {

	data := []byte{0x01, 0x0a, 0x02, 0x0a, 0x0b}
	vm := NewVM(data)
	assert.Nil(t, vm.Run())
	assert.Equal(t, byte(3), vm.stack[vm.sp])

}
