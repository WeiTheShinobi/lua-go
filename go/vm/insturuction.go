package vm

const (
	MAXARG_Bx  = 1<<18 - 1
	MAXARG_sBx = MAXARG_Bx >> 1
)

type Instruction uint32

func (ins Instruction) Opcode() int {
	return int(ins & 0x3F)
}

func (ins Instruction) ABC() (a, b, c int) {
	a = int(ins >> 6 & 0xFF)
	b = int(ins >> 14 & 0x1FF)
	c = int(ins >> 23 & 0x1FF)
	return
}

func (ins Instruction) ABx() (a, bx int) {
	a = int(ins >> 6 & 0xFF)
	bx = int(ins >> 14)
	return
}

func (ins Instruction) AsBx() (a, sbx int) {
	a, bx := ins.ABx()
	return a, bx - MAXARG_sBx
}

func (ins Instruction) Ax() int {
	return int(ins >> 6)
}

func (self Instruction) OpName() string {
	return opcodes[self.Opcode()].name
}

func (self Instruction) OpMode() byte {
	return opcodes[self.Opcode()].opMode
}

func (self Instruction) BMode() byte {
	return opcodes[self.Opcode()].argBMode
}

func (self Instruction) CMode() byte {
	return opcodes[self.Opcode()].argCMode
}
