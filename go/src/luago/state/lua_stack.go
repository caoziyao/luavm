package state

type luaStack struct {
	//slots []luaValue
	top int
}

func newLuaStack(size int) *luaStack {
	ls := &luaStack{
		top: 0,
	}
	return ls
}

func (self *luaStack) push(val luaValue) {
	//if self.top == len(self.slots) {
	//	panic("stack overflow!")
	//}
	//self.slots[self.top] = val
	//self.top++
}
