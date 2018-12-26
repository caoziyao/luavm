package vm

import . "luago/go/src/luago/api"

//
func move(i Instruction, vm LuaVM)  {

	a, b, _ := i.ABC()
	a += 1
	b += 1

	vm.Copy(b, a)
}