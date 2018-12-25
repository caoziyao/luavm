package state

func (self *luaState) PC() int {
	return self.pc
}

func (self *luaState) AddPC(n int) {
	self.pc += n
}

func (self *luaState) Fetch() uint32  {
	i := self.proto.Code[self.pc]
	self.pc += 1
	return i
}

func (self *luaState) GetConst(idx int)  {
	c := self.proto.Constants[idx]
	self.stack.push(c)
}