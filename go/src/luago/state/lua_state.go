package state

import "luago/go/src/luago/binchunk"


type luaState struct {
	stack *luaStack
	proto *binchunk.Prototype
	pc    int
}

func New() *luaState {
	ls := &luaState{
		stack: newLuaStack(20),
	}

	return ls
}
