package api

type LuaVM interface {
	LuaState
	PC() int          // 返回当前 pc
	AddPC(n int)      // 修改 pc
	Fetch() uint32    // 取出当前指令，将 pc 指向下一条指令
	GetConst(idx int) // 将指定常量推入栈顶
	GetPK(rk int)     // 将指定常量或栈值推入栈顶
}
