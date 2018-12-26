package vm

import "luago/go/src/luago/api"

const MAXARG_Bx = 1<<18 - 1       // 262143
const MAXARG_sBx = MAXARG_Bx >> 1 // 131071

type Instruction uint32

// 从指令提取操作码
func (self Instruction) Opcode() int {
	return int(self & 0x3F)
}

// 从 iABC 模式指令 提取参数
// b(9) c(9) a(8) move(6)
func (self Instruction) ABC() (a, b, c int) {

	a = int(self >> 6 & 0xFF)
	c = int(self >> 14 & 0x1FF)
	b = int(self >> 23 & 0x1FF)
	return
}

// 从 iABx 模式指令提取参数
func (self Instruction) ABx() (a, bx int) {
	a = int(self >> 6 & 0xFF)
	bx = int(self >> 14 & 0x1FF)
	return
}

// 从 iAsBx 模式指令提取参数
func (self Instruction) AsBx() (a, sbx int) {
	a, bx := self.ABx()
	return a, bx - MAXARG_sBx
}

// 从 iAx 模式指令提取参数
func (self Instruction) Ax() int {
	return int(self >> 6)
}

func (self Instruction) OpName() string  {
	return opcodes[self.Opcode()].name
}

func (self Instruction) OpMode() byte  {
	return opcodes[self.Opcode()].opMode
}

func (self Instruction) BMode() byte  {
	return opcodes[self.Opcode()].argBMode
}

func (self Instruction) CMode() byte  {
	return opcodes[self.Opcode()].argCMode
}

func (self Instruction) Execute(vm api.LuaVM)  {
	action := opcodes[self.Opcode()].action
	if action != nil {
		action(self, vm)
	} else {
		panic(self.OpName())
	}
}