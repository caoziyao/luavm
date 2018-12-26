package api

type LuaType = int

type LuaState interface {
	GetTop() int

	Copy(fromIdx, toIdx int)

	Type(idx int) LuaType

	/* access functions (stack -> Go) */
	TypeName(tp LuaType) string

	// to
	ToBoolean(idx int) bool
	//ToInteger(idx int) int64
	//ToIntegerX(idx int) (int64, bool)
	ToNumber(idx int) float64
	ToNumberX(idx int) (float64, bool)
	ToString(idx int) string
	ToStringX(idx int) (string, bool)

	/* push functions (Go -> stack) */
	//PushNil()
	PushBoolean(b bool)
}
