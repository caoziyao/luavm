package state

func (self *luaState) GetTop() int {
	return self.stack.top
}


// [-0, +0, –]
// http://www.lua.org/manual/5.3/manual.html#lua_copy
func (self *luaState) Copy(fromIdx, toIdx int)  {
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx, val)
}