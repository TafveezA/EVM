package evm

import (
	"fmt"
	"testing"
)

func TestVM(t *testing.T) {
	data := []byte{0x01, 0x0a, 0x02, 0x0a, 0x0b}
	vm := NewVM(data)
	vm.Run()
	fmt.Printf("%+v", vm.stack)
}
