package evm

type Instruction byte

const (
	InstrPushInt  Instruction = 0x0a //10
	InstrAdd      Instruction = 0x0b //1
	InstrPushByte Instruction = 0x0c //
	IntrPack      Instruction = 0x0d //
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
	data  []byte
	ip    int //instruction pointer
	stack *Stack
}

func NewVM(data []byte) *VM {
	return &VM{
		data:  data,
		ip:    0,
		stack: NewStack(128),
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
	case InstrPushByte:
		vm.stack.Push(byte(vm.data[vm.ip-1]))
	case InstrPushInt:
		vm.stack.Push(int(vm.data[vm.ip-1]))
	case InstrAdd:
		a := vm.stack.Pop().(int)
		b := vm.stack.Pop().(int)
		c := a + b
		vm.stack.Push(c)
	case InstrPack:
		n := vm.stack.Pop().(int)

	}
	return nil
}
