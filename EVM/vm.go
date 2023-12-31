package evm

import (
	"github.com/TafveezA/EVM/util"
)

type Instruction byte

const (
	InstrPushInt  Instruction = 0x0a //10
	InstrAdd      Instruction = 0x0b //1
	InstrPushByte Instruction = 0x0c //
	InstrPack     Instruction = 0x0d //
	InstrSub      Instruction = 0x0e //
	InstrStore    Instruction = 0x0f //
	InstrGet      Instruction = 0xae //
	InstrMul      Instruction = 0xea
	InstrDiv      Instruction = 0xfd
)

type Stack struct {
	data []any
	sp   int
}

func NewStack(size int) *Stack {
	return &Stack{
		data: make([]any, size),
		sp:   0,
	}

}

func (s *Stack) Push(v any) {
	s.data[s.sp] = v
	s.sp++
}

func (s *Stack) Pop() any {
	value := s.data[0]
	s.data = append(s.data[:0], s.data[1:]...)
	s.sp--
	return value
}

type VM struct {
	data          []byte
	ip            int //instruction pointer
	stack         *Stack
	contractState *State
}

func NewVM(data []byte, contractState *State) *VM {
	return &VM{
		contractState: contractState,
		data:          data,
		ip:            0,
		stack:         NewStack(128),
	}
}

func (vm *VM) Run() error {
	for {
		instr := Instruction(vm.data[vm.ip])
		//fmt.Println(instr)
		if err := vm.Exec(instr); err != nil {
			return err
		}
		vm.ip++
		if vm.ip > len(vm.data)-1 {
			break
		}
	}
	return nil
}

func (vm *VM) Exec(instr Instruction) error {
	switch instr {
	case InstrGet:
		key := vm.stack.Pop().([]byte)
		value, err := vm.contractState.Get(key)
		if err != nil {
			return err
		}
		vm.stack.Push(value)
	case InstrPushByte:
		vm.stack.Push(byte(vm.data[vm.ip-1]))
	case InstrPushInt:
		vm.stack.Push(int(vm.data[vm.ip-1]))
	case InstrSub:
		a := vm.stack.Pop().(int)
		b := vm.stack.Pop().(int)
		c := a - b
		vm.stack.Push(c)
	case InstrStore:
		var (
			key             = vm.stack.Pop().([]byte)
			value           = vm.stack.Pop()
			serializedValue []byte
		)

		switch v := value.(type) {
		case int:
			serializedValue = util.SerialzeInt64(int64(v))
		default:
			panic("TODO: unknown type")
		}

		vm.contractState.Put(key, serializedValue)

	case InstrAdd:
		a := vm.stack.Pop().(int)
		b := vm.stack.Pop().(int)
		c := a + b
		vm.stack.Push(c)

	case InstrPack:
		n := vm.stack.Pop().(int)
		//panic(n)
		b := make([]byte, n)
		for i := 0; i < n; i++ {
			b[i] = vm.stack.Pop().(byte)
		}
		vm.stack.Push(b)
	case InstrMul:
		a := vm.stack.Pop().(int)
		b := vm.stack.Pop().(int)
		c := a * b
		vm.stack.Push(c)
	case InstrDiv:
		a := vm.stack.Pop().(int)
		b := vm.stack.Pop().(int)
		c := a / b
		vm.stack.Push(c)

	}
	return nil
}
