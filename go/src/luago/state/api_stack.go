package state

func (self *luaState) GetTop() int {
	return self.stack.top
}
