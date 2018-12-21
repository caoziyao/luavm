package binchunk

type header struct {
	signature       [4]byte // 1B4C7561
	version         byte    // 5.3.4 -> 5 * 16 + 3 = 0x53， 4不用算
	format          byte    // 0
	luacData        [6]byte // 0x19930D0A1A0A
	cinSize         byte    // 0x04 在 chunk 占字节数
	sizeSize        byte    // 0x08
	instructionSize byte    // 0x04
	luaIntegerSize  byte    // 0x08
	luaNumberSize   byte    // 0x08
	luacInt         int64   // 检测大小端 0x5678 -> 0x78 56 00 00 00 00 00 00 (小端)
	luacNum         float64 // 检测浮点数格式 370.5 ->  0x00 00 00 00 00 28 77 40 (IEEE 754)
}

const (
	LUA_SIGNATURE    = "\x1bLua"
	LUAC_VERSION     = 0x53
	LUAC_FORMAT       = 0
	LUAC_DATA         =  "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSZIET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 8
	LUAC_INT          = 0x5678
	LUAC_NUM          = 370.5
)

// tag 值常量
const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x03
	TAG_INTEGER   = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x14
)

type Upvalue struct {
	Instack byte
	Idx     byte
}

type LocVar struct {
	VarName string
	StartPC uint32
	EndPC   uint32
}

/*
源文件名，起止行号，固定参数个数，
是否是 vararg 函数，寄存器数量。

指令表，常量表，upvalue 表，
子函数原型表，调试信息。

调试信息:
行号表，局部变量表，upvalue名列表
*/
type Prototype struct {
	Source          string
	LineDefined     uint32
	LastLineDefined uint32
	NumParams       byte
	IsVararg        byte
	MaxStackSize    byte
	Code            []uint32
	Constants       []interface{}
	Upvalues        []Upvalue
	Protos          []*Prototype
	LineInfo        []uint32
	LocVars         []LocVar
	UpvalueNames    []string
}

type binaryChunk struct {
	header                  // 头部
	sizeUpvalues byte       // 主函数 upvalue 数量
	mainFunc     *Prototype // 主函数原型
}

func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()
	reader.readByte() // size_upvalues
	return reader.readProto("")
}